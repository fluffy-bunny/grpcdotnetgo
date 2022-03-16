package claimsprincipal

import (
	claimsprincipalContracts "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
)

// EntryPointClaimsBuilder struct
type EntryPointClaimsBuilder struct {
	GrpcEntrypointClaimsMap map[string]*middleware_oidc.EntryPointConfig
}

// NewEntryPointClaimsBuilder ...
func NewEntryPointClaimsBuilder() *EntryPointClaimsBuilder {
	return &EntryPointClaimsBuilder{
		GrpcEntrypointClaimsMap: make(map[string]*middleware_oidc.EntryPointConfig),
	}
}

// WithGrpcEntrypointPermissionsClaimsMapOpen helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapOpen(fullMethodName string) *EntryPointClaimsBuilder {
	result := s.ensureEntry(fullMethodName)
	s.GrpcEntrypointClaimsMap[fullMethodName] = result
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapAND helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapAND(fullMethodName string, claims ...claimsprincipalContracts.Claim) *EntryPointClaimsBuilder {
	result := s.ensureEntry(fullMethodName)
	for _, claim := range claims {
		result.ClaimsConfig.AND = append(result.ClaimsConfig.AND, claim)
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapANDTYPE helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapANDTYPE(fullMethodName string, claimTypes ...string) *EntryPointClaimsBuilder {
	result := s.ensureEntry(fullMethodName)
	for _, claimType := range claimTypes {
		result.ClaimsConfig.ANDTYPE = append(result.ClaimsConfig.ANDTYPE, claimType)
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapOR helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapOR(fullMethodName string, claims ...claimsprincipalContracts.Claim) *EntryPointClaimsBuilder {
	result := s.ensureEntry(fullMethodName)
	for _, claim := range claims {
		result.ClaimsConfig.OR = append(result.ClaimsConfig.OR, claim)
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapORTYPE helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapORTYPE(fullMethodName string, claimTypes ...string) *EntryPointClaimsBuilder {
	result := s.ensureEntry(fullMethodName)
	for _, claimType := range claimTypes {
		result.ClaimsConfig.ORTYPE = append(result.ClaimsConfig.ORTYPE, claimType)
	}
	return s
}

func (s *EntryPointClaimsBuilder) ensureEntry(fullMethodName string) *middleware_oidc.EntryPointConfig {
	result, ok := s.GrpcEntrypointClaimsMap[fullMethodName]
	if !ok {
		result = &middleware_oidc.EntryPointConfig{
			FullMethodName: fullMethodName,
			ClaimsConfig:   middleware_oidc.ClaimsConfig{},
		}
		s.GrpcEntrypointClaimsMap[fullMethodName] = result
	}
	return result
}
