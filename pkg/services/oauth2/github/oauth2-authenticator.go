package github

import (
	"fmt"
	"reflect"

	"encoding/json"
	"net/http"

	contracts_oauth2 "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oauth2"
	contracts_oauth2_github "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oauth2/github"
	di "github.com/fluffy-bunny/sarulabsdi"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type (
	service struct {
		oauth2.Config
		GetOAuth2AuthenticatorConfig contracts_oauth2.GetOAuth2AuthenticatorConfig `inject:""`
		issuer                       string
	}
	gitHubResponseStruct struct {
		AvatarURL string `json:"avatar_url"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		Login     string `json:"login"`
		ID        int64  `json:"id"`
	}
	githubEmailResponse []struct {
		Email    string `json:"email"`
		Primary  bool   `json:"primary"`
		Verified bool   `json:"verified"`
	}
)

func assertImplementation() {
	var _ contracts_oauth2.IOAuth2Authenticator = (*service)(nil)
}

var reflectType = reflect.TypeOf((*service)(nil))

// AddSingletonIGithubOAuth2Authenticator registers the *service as a singleton.
func AddSingletonIGithubOAuth2Authenticator(builder *di.Builder) {
	contracts_oauth2_github.AddSingletonIGithubOAuth2Authenticator(builder, reflectType, contracts_oauth2.ReflectTypeIOAuth2Authenticator)
}
func (s *service) Ctor() {
	config := s.GetOAuth2AuthenticatorConfig()
	config.Endpoint = github.Endpoint
	s.Config = *config
}

func (s *service) GetUser(token *oauth2.Token) (*contracts_oauth2_github.User, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+token.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	githubResponse := gitHubResponseStruct{}
	err = json.NewDecoder(resp.Body).Decode(&githubResponse)
	if err != nil {
		return nil, err
	}

	gitHubEmails := githubEmailResponse{}
	req, err = http.NewRequest("GET", "https://api.github.com/user/emails", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+token.AccessToken)
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(resp.Body).Decode(&gitHubEmails)
	if err != nil {
		fmt.Println(err)
	}
	if githubResponse.Email == "" {
		githubResponse.Email = gitHubEmails[0].Email
	}
	var emails []contracts_oauth2_github.Email
	for _, v := range gitHubEmails {
		emails = append(emails, contracts_oauth2_github.Email{
			Email:    v.Email,
			Primary:  v.Primary,
			Verified: v.Verified,
		})
	}
	return &contracts_oauth2_github.User{
		Name:     githubResponse.Name,
		Email:    githubResponse.Email,
		Picture:  githubResponse.AvatarURL,
		UserName: githubResponse.Login,
		ID:       githubResponse.ID,
		Emails:   emails,
	}, nil
}
