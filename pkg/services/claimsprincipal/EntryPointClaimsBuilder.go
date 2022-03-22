package claimsprincipal

import (
	contracts_core_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
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
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimFactsMapAND(fullMethodName string, claimFacts ...*middleware_oidc.ClaimFact) *EntryPointClaimsBuilder {
	cc := s.GetClaimsConfig(fullMethodName)
	cc.WithGrpcEntrypointPermissionsClaimFactsMapAND(claimFacts...)
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapAND helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapAND(fullMethodName string, claims ...contracts_core_claimsprincipal.Claim) *EntryPointClaimsBuilder {
	for _, claim := range claims {
		s.WithGrpcEntrypointPermissionsClaimFactsMapAND(fullMethodName, &middleware_oidc.ClaimFact{
			Claim:     claim,
			Directive: middleware_oidc.ClaimTypeAndValue,
		})
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapANDTYPE helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapANDTYPE(fullMethodName string, claimTypes ...string) *EntryPointClaimsBuilder {
	for _, claimType := range claimTypes {
		s.WithGrpcEntrypointPermissionsClaimFactsMapAND(fullMethodName, &middleware_oidc.ClaimFact{
			Claim:     contracts_core_claimsprincipal.Claim{Type: claimType},
			Directive: middleware_oidc.ClaimType,
		})
	}
	return s
}

// GetClaimsConfig ...
func (s *EntryPointClaimsBuilder) GetClaimsConfig(fullMethodName string) *middleware_oidc.ClaimsConfig {
	result := s.ensureEntry(fullMethodName)
	return result.ClaimsConfig
}

// WithGrpcEntrypointPermissionsClaimFactsMapOR helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimFactsMapOR(fullMethodName string, claimFacts ...*middleware_oidc.ClaimFact) *EntryPointClaimsBuilder {
	cc := s.GetClaimsConfig(fullMethodName)
	cc.WithGrpcEntrypointPermissionsClaimFactsMapOR(claimFacts...)
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapOR helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapOR(fullMethodName string, claims ...contracts_core_claimsprincipal.Claim) *EntryPointClaimsBuilder {
	for _, claim := range claims {
		s.WithGrpcEntrypointPermissionsClaimFactsMapOR(fullMethodName, &middleware_oidc.ClaimFact{
			Claim:     claim,
			Directive: middleware_oidc.ClaimTypeAndValue,
		})
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapORTYPE helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapORTYPE(fullMethodName string, claimTypes ...string) *EntryPointClaimsBuilder {
	for _, claimType := range claimTypes {
		s.WithGrpcEntrypointPermissionsClaimFactsMapOR(fullMethodName, &middleware_oidc.ClaimFact{
			Claim:     contracts_core_claimsprincipal.Claim{Type: claimType},
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
			ClaimsConfig:   &middleware_oidc.ClaimsConfig{},
		}
		s.GrpcEntrypointClaimsMap[fullMethodName] = result
	}
	return result
}

// NewClaimFactTypeAndValue ...
func NewClaimFactTypeAndValue(claimType string, value string) *middleware_oidc.ClaimFact {
	return &middleware_oidc.ClaimFact{
		Claim: contracts_core_claimsprincipal.Claim{
			Type:  claimType,
			Value: value,
		},
		Directive: middleware_oidc.ClaimTypeAndValue,
	}
}

// NewClaimFactType ...
func NewClaimFactType(claimType string) *middleware_oidc.ClaimFact {
	return &middleware_oidc.ClaimFact{
		Claim: contracts_core_claimsprincipal.Claim{
			Type: claimType,
		},
		Directive: middleware_oidc.ClaimType,
	}
}
