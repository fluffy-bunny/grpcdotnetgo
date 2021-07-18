package mockoidcservice

import (
	servicesServiceProvider "github.com/fluffy-bunny/grpcdotnetgo/services/serviceprovider"
	"github.com/fluffy-bunny/mockoidc"
	"github.com/rs/zerolog/log"
)

type service struct {
	ServiceProvider servicesServiceProvider.IServiceProvider
	MockOIDC        *mockoidc.MockOIDC
}

func (s *service) Run() {
	s.MockOIDC, _ = mockoidc.Run()
	log.Info().Str("Issuer", s.MockOIDC.Config().Issuer).Msg("starting MockOIDC Service")
}

func (s *service) Shutdown() {
	s.MockOIDC.Shutdown()
}
