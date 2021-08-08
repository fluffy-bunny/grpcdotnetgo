package oidc

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/fluffy-bunny/go-jwt-middleware"
	"github.com/fluffy-bunny/grpcdotnetgo/middleware/oidc/introspection"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func NewGinIntrospectionValidationMiddleware(options *NewGinIntrospectionValidationMiddlewareOptions) gin.HandlerFunc {
	logger := log.New()

	logger.SetLevel(options.LogLevel)
	logger.SetOutput(options.Out)

	discoveryDocument := NewDiscoveryDocument(*options.DiscoveryURL)

	err := discoveryDocument.Initialize()
	if err != nil {
		logger.Error("Error acquiring discovery document")

		panic(err)
	}

	return func(c *gin.Context) {
		token, err := jwtmiddleware.FromAuthHeader(c.Request)
		if err != nil {
			logger.Info("Could not extract bearer token", err)

			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		err = introspection.ValidateToken(introspection.IntrospectionValidationOptions{
			Logger:           logger,
			IntrospectionURL: discoveryDocument.IntrospectionEndpoint,
			ClientID:         options.ClientID,
			ClientSecret:     options.ClientSecret,
			Token:            token,
		})
		if err != nil {
			logger.Info("Error validating token", err)
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		c.Next()
	}
}

func NewAuthenticationMiddleware(options NewJWTValidationMiddlewareOptions) *jwtmiddleware.JWTMiddleware {
	discoveryDocument := NewDiscoveryDocument(*options.DiscoveryURL)

	logger := log.New()

	logger.SetLevel(options.LogLevel)
	logger.SetOutput(options.Out)

	err := discoveryDocument.Initialize()
	if err != nil {
		logger.Error("Error acquiring discovery document")

		panic(err)
	}

	middleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: createValidationKeyGetter(discoveryDocument),
		ErrorHandler:        nil,
		Extractor:           jwtmiddleware.FromAuthHeader,
		Debug:               options.LogLevel == log.DebugLevel,
		SigningMethod:       jwt.SigningMethodRS256,
	})

	return middleware
}

func NewGinAuthenticationMiddleware(options NewJWTValidationMiddlewareOptions) gin.HandlerFunc {
	middleware := NewAuthenticationMiddleware(options)

	return func(c *gin.Context) {
		err := middleware.CheckJWT(c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		c.Next()
	}
}

func NewGrpcAuthenticationMiddleware(options NewOIDCAuthenticationOptions) func(context.Context, string) (context.Context, interface{}, error) {

	return func(ctx context.Context, fullMethodName string) (context.Context, interface{}, error) {
		return ctx, nil, nil
	}
}
