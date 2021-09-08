package di

import (
	"sync"
)

type deflist []*Def

func (e deflist) Len() int {
	return len(e)
}

func (e deflist) Less(i, j int) bool {
	return e[i].Name > e[j].Name
}

func (e deflist) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// containerCore contains the data of a Container.
// But it can not build objects on its own.
// It should be used inside a container.
type containerCore struct {
	m               sync.RWMutex
	closed          bool
	scope           string
	scopes          ScopeList
	definitions     DefMap
	parent          *containerCore
	children        map[*containerCore]struct{}
	unscopedChild   *containerCore
	objects         map[objectKey]interface{}
	lastUniqueID    int
	deleteIfNoChild bool
	dependencies    *graph
	typeDefMap      map[string]deflist
}

func (ctn *containerCore) Definitions() map[string]Def {
	return ctn.definitions.Copy()
}

func (ctn *containerCore) Scope() string {
	return ctn.scope
}

func (ctn *containerCore) Scopes() []string {
	return ctn.scopes.Copy()
}

func (ctn *containerCore) ParentScopes() []string {
	return ctn.scopes.ParentScopes(ctn.scope)
}

func (ctn *containerCore) SubScopes() []string {
	return ctn.scopes.SubScopes(ctn.scope)
}
