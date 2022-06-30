package atomic

import (
	"github.com/cristalhq/atomix"
)

// IAtomicBool interface to AtomicBool
type IAtomicBool interface {
	Store(new bool)
	Load() bool
	Swap(new bool) bool
	Toggle() bool
	String() string
	CAS(old, new bool) bool
}

// NewAtomicBool returns an instance of IAtomicBool
func NewAtomicBool(new bool) IAtomicBool {
	return atomix.NewBool(new)
}
