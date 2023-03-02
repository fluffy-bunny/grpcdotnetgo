package claimsprincipal

type EntryPointConfig struct {
	FullMethodName string     `mapstructure:"FULL_METHOD_NAME"`
	ClaimsAST      *ClaimsAST `mapstructure:"CLAIMS_CONFIG"`
}
