package oauth2

import (
	"context"

	grpc_auth "github.com/fluffy-bunny/grpcdotnetgo/pkg/go-grpc-middleware/auth"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func OAuth2UnaryServerInterceptor(oauth2Context *OAuth2Context) grpc.UnaryServerInterceptor {
	authFunc := buildAuthFunction(oauth2Context)
	return grpc_auth.UnaryServerInterceptor(authFunc)
}

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
		token, err := grpc_auth.AuthFromMD(ctx, oauth2Context.Scheme)
		if err != nil {
			// not ours
			return ctx, nil
		}
		validatedToken, err := oauth2Context.JWTValidator.ParseToken(ctx, token)

		if err != nil {
			log.Debug().Str("token", token).Msg("could not validate")
			return ctx, nil // we don't reject here, that is done in a upcomming middleware.  It is looking for that claims principal
			//	return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		} else {
			log.Debug().Str("subject", validatedToken.Token.Subject()).Msg("Validated user")
		}

		newCtx := context.WithValue(ctx, CtxClaimsPrincipalKey, validatedToken)

		return newCtx, nil
	}
}
