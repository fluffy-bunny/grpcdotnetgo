package di

import "reflect"

type TypeSet map[reflect.Type]struct{}

func NewTypeSet() TypeSet {
	return TypeSet{}
}
func (s TypeSet) Remove(rtype reflect.Type) {
	if s.Has(rtype) {
		delete(s, rtype)
	}
}
func (s TypeSet) Add(rtype reflect.Type) {
	s[rtype] = struct{}{}
}
func (s TypeSet) Has(rtype reflect.Type) bool {
	_, ok := s[rtype]
	return ok
}

// Def contains information to build and close an object inside a Container.
type Def struct {
	Build            func(ctn Container) (interface{}, error)
	Close            func(obj interface{}) error
	Name             string //[ignored] if Type is used this is overriden and hidden.
	Scope            string
	Tags             []Tag
	Type             reflect.Type //[optional] only if you want to claim that this object also implements these types.
	ImplementedTypes TypeSet
	Unshared         bool
	SafeInject       bool
	hasCtor          bool
}

// Tag can contain more specific information about a Definition.
// It is useful to find a Definition thanks to its tags instead of its name.
type Tag struct {
	Name string
	Args map[string]string
}

// DefMap is a collection of Def ordered by name.
type DefMap map[string]Def

// Copy returns a copy of the DefMap.
func (m DefMap) Copy() DefMap {
	defs := DefMap{}

	for name, def := range m {
		defs[name] = def
	}

	return defs
}
