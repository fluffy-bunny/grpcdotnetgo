package core

type ConfigOptions struct {
	Destination            interface{}
	RootConfig             []byte
	ConfigPath             string
	ApplicationEnvironment string `json:"applicationEnvironment" mapstructure:"APPLICATION_ENVIRONMENT"`
	PrettyLog              bool   `json:"prettyLog" mapstructure:"PRETTY_LOG"`
	LogLevel               string `json:"logLevel" mapstructure:"LOG_LEVEL"`
}
