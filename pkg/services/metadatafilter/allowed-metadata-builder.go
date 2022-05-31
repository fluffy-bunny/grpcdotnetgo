package metadatafilter

import (
	"strings"

	"github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
)

// EntryPointAllowedMetadataMapBuilder struct
type EntryPointAllowedMetadataMapBuilder struct {
	EntryPointAllowedMetadataMap map[string]*hashset.StringSet
}

// NewEntryPointAllowedMetadataMapBuilder ...
func NewEntryPointAllowedMetadataMapBuilder() *EntryPointAllowedMetadataMapBuilder {
	return &EntryPointAllowedMetadataMapBuilder{
		EntryPointAllowedMetadataMap: make(map[string]*hashset.StringSet),
	}
}

// WithAllowedMetadataHeader helper to add a single entrypoint config
func (s *EntryPointAllowedMetadataMapBuilder) WithAllowedMetadataHeader(fullMethodName string, headers ...string) *EntryPointAllowedMetadataMapBuilder {
	_, ok := s.EntryPointAllowedMetadataMap[fullMethodName]
	if !ok {
		s.EntryPointAllowedMetadataMap[fullMethodName] = hashset.NewStringSet()
	}
	for _, header := range headers {
		s.EntryPointAllowedMetadataMap[fullMethodName].Add(strings.ToLower(header))
	}
	return s
}
