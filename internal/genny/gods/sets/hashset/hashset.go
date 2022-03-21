// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

package hashset

import (
	"fmt"
	"strings"

	"github.com/cheekybits/genny/generic"
)

// KeyType  ...
type KeyType generic.Type

// KeyTypeSet holds elements in go's native map
type KeyTypeSet struct {
	items map[KeyType]struct{}
}

var itemKeyTypeExists = struct{}{}

// NewKeyTypeSet instantiates a new empty set and adds the passed values, if any, to the set
func NewKeyTypeSet(values ...KeyType) *KeyTypeSet {
	set := &KeyTypeSet{items: make(map[KeyType]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *KeyTypeSet) Add(items ...KeyType) {
	for _, item := range items {
		set.items[item] = itemKeyTypeExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *KeyTypeSet) Remove(items ...KeyType) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *KeyTypeSet) Contains(items ...KeyType) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *KeyTypeSet) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *KeyTypeSet) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *KeyTypeSet) Clear() {
	set.items = make(map[KeyType]struct{})
}

// Values returns all items in the set.
func (set *KeyTypeSet) Values() []KeyType {
	values := make([]KeyType, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *KeyTypeSet) String() string {
	str := "KeyTypeSet\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}
