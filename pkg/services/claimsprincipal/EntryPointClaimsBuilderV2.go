package claimsprincipal

import (
	contracts_claimsprincipal "github.com/fluffy-bunny/grpcdotnetgo/pkg/contracts/claimsprincipal"
)

// EntryPointClaimsBuilderV2 struct
type EntryPointClaimsBuilderV2 struct {
	GrpcEntrypointClaimsMap map[string]*EntryPointConfig
}

// NewEntryPointClaimsBuilderV2 ...
func NewEntryPointClaimsBuilderV2() *EntryPointClaimsBuilderV2 {
	return &EntryPointClaimsBuilderV2{
		GrpcEntrypointClaimsMap: make(map[string]*EntryPointConfig),
	}
}

// WithGrpcEntrypointPermissionsClaimsMapOpen helper to add a single entrypoint config
func (s *EntryPointClaimsBuilderV2) WithGrpcEntrypointPermissionsClaimsMapOpen(fullMethodName string) *EntryPointClaimsBuilderV2 {
	s.ensureEntry(fullMethodName)
	return s
}

func (s *EntryPointClaimsBuilderV2) WithGrpcEntrypointClams(fullMethodName string, claims ...contracts_claimsprincipal.Claim) *EntryPointClaimsBuilderV2 {
	ast := s.GetClaimsAST(fullMethodName)
	ast.Claims = append(ast.Claims, claims...)
	return s
}

// GetClaimsAST ...
func (s *EntryPointClaimsBuilderV2) GetClaimsAST(fullMethodName string) *ClaimsAST {
	result := s.ensureEntry(fullMethodName)
	return result.ClaimsAST
}

func (s *EntryPointClaimsBuilderV2) ensureEntry(fullMethodName string) *EntryPointConfig {
	result, ok := s.GrpcEntrypointClaimsMap[fullMethodName]
	if !ok {
		result = &EntryPointConfig{
			FullMethodName: fullMethodName,
			ClaimsAST:      &ClaimsAST{},
		}
		s.GrpcEntrypointClaimsMap[fullMethodName] = result
	}
	return result
}
