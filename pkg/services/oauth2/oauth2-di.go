package oauth2

import (
	"reflect"

	contractsOAuth2 "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/oauth2"
	di "github.com/fluffy-bunny/sarulabsdi"
)

// AddOauth2Service adds service to the DI container
func AddOauth2Service(builder *di.Builder) {
	contractsOAuth2.AddScopedIOauth2(builder, reflect.TypeOf(&service{}))
}
