package metadatafilter

import "strings"

// EntryPointAllowedMetadataMapBuilder struct
type EntryPointAllowedMetadataMapBuilder struct {
	EntryPointAllowedMetadataMap map[string]map[string]bool
}

// NewEntryPointAllowedMetadataMapBuilder ...
func NewEntryPointAllowedMetadataMapBuilder() *EntryPointAllowedMetadataMapBuilder {
	return &EntryPointAllowedMetadataMapBuilder{
		EntryPointAllowedMetadataMap: make(map[string]map[string]bool),
	}
}

// WithAllowedMetadataHeader helper to add a single entrypoint config
func (s *EntryPointAllowedMetadataMapBuilder) WithAllowedMetadataHeader(fullMethodName string, headers ...string) *EntryPointAllowedMetadataMapBuilder {
	_, ok := s.EntryPointAllowedMetadataMap[fullMethodName]
	if !ok {
		s.EntryPointAllowedMetadataMap[fullMethodName] = make(map[string]bool)
	}
	for _, header := range headers {
		s.EntryPointAllowedMetadataMap[fullMethodName][strings.ToLower(header)] = true
	}
	return s
}
