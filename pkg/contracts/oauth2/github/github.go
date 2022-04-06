package github

import (
	contracts_oauth2 "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oauth2"
	"golang.org/x/oauth2"
)

//go:generate genny -pkg $GOPACKAGE -in=../../../../genny/sarulabsdi/interface-types.go -out=gen-$GOFILE gen "InterfaceType=IGithubOAuth2Authenticator"

//go:generate mockgen -package=$GOPACKAGE -destination=../../../mocks/oauth2/$GOPACKAGE/mock_$GOFILE   github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oauth2/$GOPACKAGE IGithubOAuth2Authenticator

type (
	// User ...
	User struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Picture  string `json:"picture"`
		UserName string `json:"userName"`
		ID       int64  `json:"id"`
		Emails []Email `json:"emails"`
	}
	Email struct {
		Email    string `json:"email"`
		Primary  bool   `json:"primary"`
		Verified bool   `json:"verified"`
	}
	// IGithubOAuth2Authenticator ...
	IGithubOAuth2Authenticator interface {
		contracts_oauth2.IOAuth2Authenticator
		GetUser(token *oauth2.Token) (*User, error)
	}
)
