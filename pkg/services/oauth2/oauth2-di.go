package oauth2

import (
	"reflect"

	contractsOAuth2 "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oauth2"
	di "github.com/fluffy-bunny/sarulabsdi"
	"github.com/rs/zerolog/log"
)

// AddOauth2Service adds service to the DI container
func AddOauth2Service(builder *di.Builder) {
	log.Info().
		Msg("IoC: AddOauth2Service")
	contractsOAuth2.AddScopedIOauth2(builder, reflect.TypeOf(&service{}))
}
