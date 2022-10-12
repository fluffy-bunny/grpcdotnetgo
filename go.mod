module github.com/fluffy-bunny/grpcdotnetgo

go 1.18

require (
	github.com/ReneKroon/ttlcache/v2 v2.11.0
	github.com/bamzi/jobrunner v1.0.0
	github.com/cheekybits/genny v1.0.0
	github.com/coreos/go-oidc/v3 v3.2.0
	github.com/fatih/structs v1.1.0
	github.com/fluffy-bunny/mockoidc v0.0.0-20210902160455-4c83c82b8422
	github.com/fluffy-bunny/sarulabsdi v0.1.63
	github.com/fluffy-bunny/viperEx v0.0.26
	github.com/form3tech-oss/jwt-go v3.2.5+incompatible
	github.com/gogo/googleapis v1.4.1
	github.com/gogo/protobuf v1.3.2
	github.com/gogo/status v1.1.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/gorilla/securecookie v1.1.1
	github.com/gorilla/sessions v1.2.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/jedib0t/go-pretty/v6 v6.3.3
	github.com/jinzhu/copier v0.3.5
	github.com/jnewmano/grpc-json-proxy v0.0.6
	github.com/labstack/echo-contrib v0.13.0
	github.com/labstack/echo/v4 v4.9.0
	github.com/lestrrat-go/jwx v1.2.25
	github.com/pkg/errors v0.9.1
	github.com/reugn/async v0.0.0-20200819063434-15e5b3951cd7
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/xid v1.4.0
	github.com/rs/zerolog v1.27.0
	github.com/satori/go.uuid v1.2.0
	github.com/segmentio/ksuid v1.0.4
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.6.0
	github.com/spf13/viper v1.12.0
	github.com/stretchr/testify v1.7.4
	github.com/tkuchiki/parsetime v0.3.0
	github.com/ziflex/lecho v1.2.0
	golang.org/x/oauth2 v0.0.0-20220411215720-9780585627b5
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/square/go-jose.v2 v2.6.0

)

require (
	cloud.google.com/go/compute v1.6.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.0-20210816181553-5444fa50b93d // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/context v1.1.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/labstack/echo v3.3.10+incompatible // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/lestrrat-go/backoff/v2 v2.0.8 // indirect
	github.com/lestrrat-go/blackmagic v1.0.0 // indirect
	github.com/lestrrat-go/httpcc v1.0.1 // indirect
	github.com/lestrrat-go/iter v1.0.1 // indirect
	github.com/lestrrat-go/option v1.0.0 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	github.com/tkuchiki/go-timezone v0.2.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/net v0.0.0-20220728030405-41545e8bf201 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220722155302-e5dcc9cfc0b9 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace github.com/fluffy-bunny/go-jwt-middleware => ../go-jwt-middleware

//replace github.com/grpc-ecosystem/go-grpc-middleware => ../go-grpc-middleware

//replace github.com/fluffy-bunny/mockoidc => ../mockoidc

//replace github.com/fluffy-bunny/sarulabsdi => ../sarulabsdi
