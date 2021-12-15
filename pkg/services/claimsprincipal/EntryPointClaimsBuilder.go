package claimsprincipal

import (
	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
)

// EntryPointClaimsBuilder struct
type EntryPointClaimsBuilder struct {
	GrpcEntrypointClaimsMap map[string]middleware_oidc.EntryPointConfig
}

// NewEntryPointClaimsBuilder ...
func NewEntryPointClaimsBuilder() *EntryPointClaimsBuilder {
	return &EntryPointClaimsBuilder{
		GrpcEntrypointClaimsMap: make(map[string]middleware_oidc.EntryPointConfig),
	}
}

// WithGrpcEntrypointPermissionsClaimsMapOpen helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapOpen(fullMethodName string) *EntryPointClaimsBuilder {
	result := middleware_oidc.EntryPointConfig{
		FullMethodName: fullMethodName,
		ClaimsConfig:   middleware_oidc.ClaimsConfig{},
	}

	s.GrpcEntrypointClaimsMap[fullMethodName] = result
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapAND helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapAND(fullMethodName string, claims ...claimsprincipalContracts.Claim) *EntryPointClaimsBuilder {
	result := middleware_oidc.EntryPointConfig{
		FullMethodName: fullMethodName,
		ClaimsConfig:   middleware_oidc.ClaimsConfig{},
	}
	for _, claim := range claims {
		result.ClaimsConfig.AND = append(result.ClaimsConfig.AND, claim)
	}
	s.GrpcEntrypointClaimsMap[fullMethodName] = result
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapOR helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapOR(fullMethodName string, claims ...claimsprincipalContracts.Claim) *EntryPointClaimsBuilder {
	result := middleware_oidc.EntryPointConfig{
		FullMethodName: fullMethodName,
		ClaimsConfig:   middleware_oidc.ClaimsConfig{},
	}
	for _, claim := range claims {
		result.ClaimsConfig.OR = append(result.ClaimsConfig.OR, claim)
	}
	s.GrpcEntrypointClaimsMap[fullMethodName] = result
	return s
}
