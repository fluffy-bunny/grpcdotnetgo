package metadatafilter

import (
	"github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	wellknown "github.com/fluffy-bunny/grpcdotnetgo/pkg/wellknown"
)

// NewHeaderSet ...
func NewHeaderSet() *hashset.StringSet {
	set := hashset.NewStringSet(wellknown.MetaDataFilter...)
	return set
}
