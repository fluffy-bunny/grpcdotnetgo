package cookies

import (
	contracts_cookies "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/cookies"

	"encoding/base64"
	"net/http"
	"reflect"
	"time"

	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/gorilla/securecookie"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type (
	service struct {
		SecureCookieConfigAccessor contracts_cookies.SecureCookieConfigAccessor `inject:" "`
		secureCookie               *securecookie.SecureCookie
	}
)

func assertImplementation() {
	var _ contracts_cookies.ISecureCookie = (*service)(nil)
}

func (s *service) Ctor() {
	config := s.SecureCookieConfigAccessor()
	hashKey, err := base64.StdEncoding.DecodeString(config.SecureCookieHashKey)
	if err != nil {
		panic(err)
	}
	encryptionKey, err := base64.StdEncoding.DecodeString(config.SecureCookieEncryptionKey)
	if err != nil {
		panic(err)
	}
	s.secureCookie = securecookie.New(hashKey, encryptionKey)
}

// AddSingletonISecureCookie ...
func AddSingletonISecureCookie(builder *di.Builder) {
	log.Info().Str("DI", "ISecureCookie").Send()
	contracts_cookies.AddSingletonISecureCookie(builder, reflect.TypeOf((*service)(nil)))
}

type valueContainer struct {
	Value string
}

func (s *service) newCookie(name string, value string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Path = "/"
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteStrictMode
	return cookie
}
func (s *service) SetCookieValue(c echo.Context, name string, value string, expires time.Time) error {
	encoded, err := s.secureCookie.Encode(name, &valueContainer{
		Value: value,
	})
	if err != nil {
		return err
	}

	cookie := s.newCookie(name, encoded)
	cookie.Expires = expires
	c.SetCookie(cookie)
	return nil
}
func (s *service) GetCookieValue(c echo.Context, name string) (string, error) {
	cookie, err := c.Cookie(name)
	if err != nil {
		return "", err
	}
	var value = &valueContainer{}

	err = s.secureCookie.Decode(name, cookie.Value, &value)
	if err != nil {
		return "", err
	}
	return value.Value, nil
}
func (s *service) DeleteCookie(c echo.Context, name string) error {
	cookie := s.newCookie(name, "")
	cookie.Expires = time.Now().Add(-24 * 365 * time.Hour)
	cookie.MaxAge = -1
	c.SetCookie(cookie)
	return nil
}
func (s *service) RefreshCookie(c echo.Context, name string, durration time.Duration) error {
	cookie, err := c.Cookie(name)
	if err != nil {
		return err
	}
	cookie = s.newCookie(name, cookie.Value)
	cookie.Expires = time.Now().Add(durration)

	c.SetCookie(cookie)
	return nil
}
