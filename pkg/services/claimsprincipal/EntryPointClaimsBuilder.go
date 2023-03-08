package claimsprincipal

import (
	contracts_core_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
	middleware_oidc "github.com/fluffy-bunny/grpcdotnetgo/pkg/middleware/oidc"
	services_claimfact "github.com/fluffy-bunny/grpcdotnetgo/pkg/services/claimfact"
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

// AddMetaData ...
func (s *EntryPointClaimsBuilder) AddMetaData(fullMethodName string, metaData map[string]interface{}) *EntryPointClaimsBuilder {
	entry, ok := s.GrpcEntrypointClaimsMap[fullMethodName]
	if !ok {
		panic("EntryPointClaimsBuilder.AddMetaData: entry not found, it must be created before metadata can be added")
	}
	entry.MetaData = metaData
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapOpen helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapOpen(fullMethodName string) *EntryPointClaimsBuilder {
	s.ensureEntry(fullMethodName)
	return s
}

// WithGrpcEntrypointPermissionsClaimFactsMapAND helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimFactsMapAND(fullMethodName string, claimFacts ...*services_claimfact.ClaimFact) *EntryPointClaimsBuilder {
	cc := s.GetClaimsConfig(fullMethodName)
	cc.WithGrpcEntrypointPermissionsClaimFactsMapAND(claimFacts...)
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapAND helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapAND(fullMethodName string, claims ...contracts_core_claimsprincipal.Claim) *EntryPointClaimsBuilder {
	for _, claim := range claims {
		s.WithGrpcEntrypointPermissionsClaimFactsMapAND(fullMethodName, &services_claimfact.ClaimFact{
			Claim:     claim,
			Directive: services_claimfact.ClaimTypeAndValue,
		})
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapANDTYPE helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapANDTYPE(fullMethodName string, claimTypes ...string) *EntryPointClaimsBuilder {
	for _, claimType := range claimTypes {
		s.WithGrpcEntrypointPermissionsClaimFactsMapAND(fullMethodName, &services_claimfact.ClaimFact{
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
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimFactsMapOR(fullMethodName string, claimFacts ...*services_claimfact.ClaimFact) *EntryPointClaimsBuilder {
	cc := s.GetClaimsConfig(fullMethodName)
	cc.WithGrpcEntrypointPermissionsClaimFactsMapOR(claimFacts...)
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapOR helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapOR(fullMethodName string, claims ...contracts_core_claimsprincipal.Claim) *EntryPointClaimsBuilder {
	for _, claim := range claims {
		s.WithGrpcEntrypointPermissionsClaimFactsMapOR(fullMethodName, &services_claimfact.ClaimFact{
			Claim:     claim,
			Directive: services_claimfact.ClaimTypeAndValue,
		})
	}
	return s
}

// WithGrpcEntrypointPermissionsClaimsMapORTYPE helper to add a single entrypoint config
func (s *EntryPointClaimsBuilder) WithGrpcEntrypointPermissionsClaimsMapORTYPE(fullMethodName string, claimTypes ...string) *EntryPointClaimsBuilder {
	for _, claimType := range claimTypes {
		s.WithGrpcEntrypointPermissionsClaimFactsMapOR(fullMethodName, &services_claimfact.ClaimFact{
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
