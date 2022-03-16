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
	s.ensureEntry(fullMethodName)
	return s
}

// WithGrpcEntrypointPermissionsClaimFactsMapAND helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimFactsMapAND(fullMethodName string, claimFacts ...middleware_oidc.ClaimFact) *EntryPointClaimsBuilder {
	result := s.ensureEntry(fullMethodName)
	for _, claimFact := range claimFacts {
		result.ClaimsConfig.AND = append(result.ClaimsConfig.AND, claimFact)
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapAND helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapAND(fullMethodName string, claims ...claimsprincipalContracts.Claim) *EntryPointClaimsBuilder {
	for _, claim := range claims {
		s.WithGrpcEntrypointPermissionsClaimFactsMapAND(fullMethodName, middleware_oidc.ClaimFact{
			Claim:     claim,
			Directive: middleware_oidc.ClaimTypeAndValue,
		})
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapANDTYPE helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapANDTYPE(fullMethodName string, claimTypes ...string) *EntryPointClaimsBuilder {
	for _, claimType := range claimTypes {
		s.WithGrpcEntrypointPermissionsClaimFactsMapAND(fullMethodName, middleware_oidc.ClaimFact{
			Claim:     claimsprincipalContracts.Claim{Type: claimType},
			Directive: middleware_oidc.ClaimType,
		})
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimFactsMapOR helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimFactsMapOR(fullMethodName string, claimFacts ...middleware_oidc.ClaimFact) *EntryPointClaimsBuilder {
	result := s.ensureEntry(fullMethodName)
	for _, claimFact := range claimFacts {
		result.ClaimsConfig.OR = append(result.ClaimsConfig.OR, claimFact)
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapOR helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapOR(fullMethodName string, claims ...claimsprincipalContracts.Claim) *EntryPointClaimsBuilder {
	for _, claim := range claims {
		s.WithGrpcEntrypointPermissionsClaimFactsMapOR(fullMethodName, middleware_oidc.ClaimFact{
			Claim:     claim,
			Directive: middleware_oidc.ClaimTypeAndValue,
		})
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapORTYPE helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapORTYPE(fullMethodName string, claimTypes ...string) *EntryPointClaimsBuilder {
	for _, claimType := range claimTypes {
		s.WithGrpcEntrypointPermissionsClaimFactsMapOR(fullMethodName, middleware_oidc.ClaimFact{
			Claim:     claimsprincipalContracts.Claim{Type: claimType},
			Directive: middleware_oidc.ClaimType,
		})
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
