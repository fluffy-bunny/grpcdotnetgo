package oidc

import (
	"errors"
	"time"

	jwtgoForm3Tech "github.com/form3tech-oss/jwt-go"
	"github.com/golang-jwt/jwt"
)

func getPemCert(JWKSResponse *JSONWebKeyResponse, token *jwtgoForm3Tech.Token) (string, error) {
	cert := ""

	for k := range JWKSResponse.Keys {
		if token.Header["kid"] == JWKSResponse.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + JWKSResponse.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("unable to find appropriate key")

		return cert, err
	}

	return cert, nil
}

func createValidationKeyGetter(doc *DiscoveryDocument) func(token *jwtgoForm3Tech.Token) (interface{}, error) {
	return func(token *jwtgoForm3Tech.Token) (interface{}, error) {
		//checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(, true)
		//if !checkAud {
		//	return token, errors.New("invalid audience")
		//}

		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(doc.Issuer, true)
		if !checkIss {
			return token, errors.New("invalid issuer")
		}

		checkTime := token.Claims.(jwt.MapClaims).VerifyIssuedAt(time.Now().Unix(), true)
		if !checkTime {
			return token, errors.New("invalid issued at")
		}

		cert, err := getPemCert(doc.KeyResponse, token)
		if err != nil {
			panic(err.Error())
		}

		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))

		return result, nil
	}
}
