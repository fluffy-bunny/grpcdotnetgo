package internal

type Config struct {
	Environment string `mapstructure:"APPLICATION_ENVIRONMENT"`
	Mode        string `mapstructure:"MODE"`
}

// ConfigDefaultYaml default yaml
var ConfigDefaultYaml = []byte(`
APPLICATION_ENVIRONMENT: in-environment
`)
