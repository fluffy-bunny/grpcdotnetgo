package oidc

import (
	"errors"

	jwtgoForm3Tech "github.com/form3tech-oss/jwt-go"
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
