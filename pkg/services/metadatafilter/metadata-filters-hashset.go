package metadatafilter

import (
	sets "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	wellknown "github.com/fluffy-bunny/grpcdotnetgo/pkg/wellknown"
)

// NewHeaderSet ...
func NewHeaderSet() *sets.StringSet {
	set := sets.NewStringSet(wellknown.MetaDataFilter...)
	return set
}
