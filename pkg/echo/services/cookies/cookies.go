package cookies

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	contracts_contextaccessor "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/contextaccessor"
	contracts_cookies "github.com/fluffy-bunny/grpcdotnetgo/pkg/echo/contracts/cookies"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/gorilla/securecookie"
	"github.com/rs/zerolog/log"
)

type (
	service struct {
		EchoContextAccessor        contracts_contextaccessor.IEchoContextAccessor `inject:""`
		SecureCookieConfigAccessor contracts_cookies.SecureCookieConfigAccessor   `inject:" "`
		secureCookie               *securecookie.SecureCookie
	}

	chunkMetaData struct {
		NumberOfChunks int    `json:"numberOfChunks"`
		Value          string `json:"value"`
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

// AddScopedISecureCookie ...
func AddScopedISecureCookie(builder *di.Builder) {
	log.Info().Str("DI", "ISecureCookie - SCOPED").Send()
	contracts_cookies.AddScopedISecureCookie(builder, reflect.TypeOf((*service)(nil)))
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
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
func (s *service) _setCookieValue(name string, value string, expires time.Time) error {
	c := s.EchoContextAccessor.GetContext()
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

func (s *service) SetCookieValue(name string, value string, expires time.Time) error {

	var chunks []string
	chunkSize := 1024
	if len(value) > chunkSize {
		for i := 0; i < len(value); i += chunkSize {
			chunk := value[i:min(i+chunkSize, len(value))]
			chunks = append(chunks, chunk)
		}
		jsonMD, _ := json.Marshal(&chunkMetaData{
			NumberOfChunks: len(chunks),
		})
		var cookieNames = []string{}
		var onError = func() {
			for _, n := range cookieNames {
				s.DeleteCookie(n)
			}
		}

		err := s._setCookieValue(name, string(jsonMD), expires) // store the number of chunks in the main cookie
		if err != nil {
			return err
		}
		cookieNames = append(cookieNames, name)
		for i, chunk := range chunks {
			chunkName := fmt.Sprintf("%s_%d", name, i)
			err := s._setCookieValue(chunkName, chunk, expires)
			if err != nil {
				onError()
				return err
			}
			cookieNames = append(cookieNames, chunkName)
		}
	} else {
		jsonMD, _ := json.Marshal(&chunkMetaData{
			NumberOfChunks: 0,
			Value:          string(value),
		})
		err := s._setCookieValue(name, string(jsonMD), expires) // store the number of chunks in the main cookie
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *service) GetCookieValue(name string) (string, error) {
	c := s.EchoContextAccessor.GetContext()
	cookie, err := c.Cookie(name)
	if err != nil {
		return "", err
	}
	var value = &valueContainer{}

	err = s.secureCookie.Decode(name, cookie.Value, value)
	if err != nil {
		return "", err
	}
	var metaData = &chunkMetaData{}
	err = json.Unmarshal([]byte(value.Value), &metaData)
	if err != nil {
		return "", err
	}
	if metaData.NumberOfChunks == 0 {
		return value.Value, nil
	}
	sbOri := strings.Builder{}
	for i := 0; i < metaData.NumberOfChunks; i++ {
		chunkName := fmt.Sprintf("%s_%d", name, i)
		chunkCookie, err := c.Cookie(chunkName)
		if err != nil {
			return "", err
		}
		var value = &valueContainer{}
		err = s.secureCookie.Decode(chunkName, chunkCookie.Value, value)
		if err != nil {
			return "", err
		}
		sbOri.WriteString(value.Value)
	}
	ori := sbOri.String()
	return ori, nil
}
func (s *service) _delete(name string) error {
	c := s.EchoContextAccessor.GetContext()
	cookie := s.newCookie(name, "")
	cookie.Expires = time.Now().Add(-24 * 365 * time.Hour)
	cookie.MaxAge = -1
	c.SetCookie(cookie)
	return nil
}
func (s *service) DeleteCookie(name string) error {
	c := s.EchoContextAccessor.GetContext()
	cookie := s.newCookie(name, "")

	cookie, err := c.Cookie(name)
	if err != nil {
		return err
	}
	s._delete(name) // delete the main cookie no matter what
	var value = &valueContainer{}

	err = s.secureCookie.Decode(name, cookie.Value, value)
	if err != nil {
		return err
	}
	var metaData = &chunkMetaData{}
	err = json.Unmarshal([]byte(value.Value), &metaData)
	if err != nil {
		return err
	}

	for i := 0; i < metaData.NumberOfChunks; i++ {
		chunkName := fmt.Sprintf("%s_%d", name, i)
		s._delete(chunkName)
	}

	return nil
}
func (s *service) _refresh(cookie *http.Cookie, duration time.Duration) {
	c := s.EchoContextAccessor.GetContext()
	newCookie := s.newCookie(cookie.Name, cookie.Value)
	newCookie.Expires = time.Now().Add(duration)
	c.SetCookie(newCookie)
}
func (s *service) RefreshCookie(name string, duration time.Duration) error {
	c := s.EchoContextAccessor.GetContext()
	cookie, err := c.Cookie(name)
	if err != nil {
		return err
	}
	var value = &valueContainer{}

	err = s.secureCookie.Decode(name, cookie.Value, value)
	if err != nil {
		return err
	}
	var metaData = &chunkMetaData{}
	err = json.Unmarshal([]byte(value.Value), &metaData)
	if err != nil {
		return err
	}

	s._refresh(cookie, duration)
	for i := 0; i < metaData.NumberOfChunks; i++ {
		chunkName := fmt.Sprintf("%s_%d", name, i)
		chunkCookie, err := c.Cookie(chunkName)
		if err != nil {
			return err
		}
		s._refresh(chunkCookie, duration)
	}

	return nil
}
