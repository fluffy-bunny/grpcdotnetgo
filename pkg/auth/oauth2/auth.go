package oauth2

import (
	"context"

	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	grpc_auth "github.com/fluffy-bunny/grpcdotnetgo/pkg/go-grpc-middleware/auth"
	dicontext "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/dicontext"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// OAuth2UnaryServerInterceptor ...
func OAuth2UnaryServerInterceptor(oauth2Context *OAuth2Context) grpc.UnaryServerInterceptor {
	authFunc := buildAuthFunction(oauth2Context)
	return grpc_auth.UnaryServerInterceptor(authFunc)
}

// BuildOAuth2Context ...
func BuildOAuth2Context(issuer string, JWKSURL string, config *GrpcFuncAuthConfig) (*OAuth2Context, error) {
	oauth2DiscoveryOptions := OAuth2DiscoveryOptions{
		JWKSURL: JWKSURL,
	}
	disco, err := newOAuth2Document(&oauth2DiscoveryOptions)
	disco.Issuer = issuer

	if err != nil {
		log.Fatal().Err(err).Str("jwksUrl", JWKSURL).Msg("Cound not newOAuth2Document")
		return nil, err
	}
	err = disco.initialize()
	if err != nil {
		log.Fatal().Err(err).Str("jwksUrl", JWKSURL).Msg("Cound not initialize discoveryDocument")
		return nil, err
	}
	log.Log().Object("disco", disco)
	jwtValidatorOptions := JWTValidatorOptions{
		ClockSkewMinutes: config.ClockSkewMinutes,
		OAuth2Document:   disco,
	}
	jwtValidator := NewJWTValidator(&jwtValidatorOptions)
	oauth2Context := OAuth2Context{
		OAuth2Document: disco,
		JWTValidator:   jwtValidator,
		Scheme:         config.ExpectedScheme,
		Config:         config,
	}
	return &oauth2Context, nil
}

// BuildOpenIdConnectContext ...
func BuildOpenIdConnectContext(config *GrpcFuncAuthConfig) (*OAuth2Context, error) {
	discoveryDocumentOptions := DiscoveryDocumentOptions{
		Authority: config.Authority,
	}
	disco, err := newDiscoveryDocument(&discoveryDocumentOptions)
	if err != nil {
		log.Fatal().Err(err).Str("authority", config.Authority).Msg("Cound not NewDiscoveryDocument")
		return nil, err
	}
	err = disco.initialize()
	if err != nil {
		log.Fatal().Err(err).Str("authority", config.Authority).Msg("Cound not initialize discoveryDocument")
		return nil, err
	}
	log.Log().Object("disco", disco)
	oauth2Context, err := BuildOAuth2Context(disco.Issuer, disco.JWKSURL, config)
	return oauth2Context, nil
}
func buildAuthFunction(oauth2Context *OAuth2Context) func(ctx context.Context, fullMethodName string) (context.Context, error) {
	return func(ctx context.Context, fullMethodName string) (context.Context, error) {
		requestContainer := dicontext.GetRequestContainer(ctx)
		claimsPrincipal := claimsprincipalContracts.GetIClaimsPrincipalFromContainer(requestContainer)

		token, err := grpc_auth.AuthFromMD(ctx, oauth2Context.Scheme)
		if err != nil {
			emptyPrincipal := oauth2Context.JWTValidator.NewEmptyClaimsPrincipal()
			// not ours
			newCtx := context.WithValue(ctx, CtxClaimsPrincipalKey, emptyPrincipal)
			return newCtx, nil
		}
		validatedToken, err := oauth2Context.JWTValidator.ParseToken(ctx, token)
		var newCtx context.Context
		if err != nil {
			log.Debug().Str("token", token).Msg("could not validate, returning empty claims principal")
			// make an empty one
			validatedToken = oauth2Context.JWTValidator.NewEmptyClaimsPrincipal()
		} else {
			log.Debug().Str("subject", validatedToken.Token.Subject()).Msg("Validated user")
		}
		for _, c := range validatedToken.Claims {
			claimsPrincipal.AddClaim(c)
		}

		newCtx = context.WithValue(ctx, CtxClaimsPrincipalKey, validatedToken)

		return newCtx, nil
	}
}
