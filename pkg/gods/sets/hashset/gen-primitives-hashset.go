// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

package sets

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StringSet holds elements in go's native map
type StringSet struct {
	items map[string]struct{}
}

var itemStringExists = struct{}{}

// NewStringSet instantiates a new empty set and adds the passed values, if any, to the set
func NewStringSet(values ...string) *StringSet {
	set := &StringSet{items: make(map[string]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *StringSet) Add(items ...string) {
	for _, item := range items {
		set.items[item] = itemStringExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *StringSet) Remove(items ...string) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *StringSet) Contains(items ...string) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *StringSet) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *StringSet) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *StringSet) Clear() {
	set.items = make(map[string]struct{})
}

// Values returns all items in the set.
func (set *StringSet) Values() []string {
	values := make([]string, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *StringSet) String() string {
	str := "StringSet\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

// ToJSON outputs the JSON representation of the set.
func (set *StringSet) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *StringSet) FromJSON(data []byte) error {
	var elements []string
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

// IntSet holds elements in go's native map
type IntSet struct {
	items map[int]struct{}
}

var itemIntExists = struct{}{}

// NewIntSet instantiates a new empty set and adds the passed values, if any, to the set
func NewIntSet(values ...int) *IntSet {
	set := &IntSet{items: make(map[int]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *IntSet) Add(items ...int) {
	for _, item := range items {
		set.items[item] = itemIntExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *IntSet) Remove(items ...int) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *IntSet) Contains(items ...int) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *IntSet) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *IntSet) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *IntSet) Clear() {
	set.items = make(map[int]struct{})
}

// Values returns all items in the set.
func (set *IntSet) Values() []int {
	values := make([]int, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *IntSet) String() string {
	str := "IntSet\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

// ToJSON outputs the JSON representation of the set.
func (set *IntSet) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *IntSet) FromJSON(data []byte) error {
	var elements []int
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

// Int32Set holds elements in go's native map
type Int32Set struct {
	items map[int32]struct{}
}

var itemInt32Exists = struct{}{}

// NewInt32Set instantiates a new empty set and adds the passed values, if any, to the set
func NewInt32Set(values ...int32) *Int32Set {
	set := &Int32Set{items: make(map[int32]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Int32Set) Add(items ...int32) {
	for _, item := range items {
		set.items[item] = itemInt32Exists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Int32Set) Remove(items ...int32) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Int32Set) Contains(items ...int32) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Int32Set) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Int32Set) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Int32Set) Clear() {
	set.items = make(map[int32]struct{})
}

// Values returns all items in the set.
func (set *Int32Set) Values() []int32 {
	values := make([]int32, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Int32Set) String() string {
	str := "Int32Set\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

// ToJSON outputs the JSON representation of the set.
func (set *Int32Set) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Int32Set) FromJSON(data []byte) error {
	var elements []int32
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

// Int64Set holds elements in go's native map
type Int64Set struct {
	items map[int64]struct{}
}

var itemInt64Exists = struct{}{}

// NewInt64Set instantiates a new empty set and adds the passed values, if any, to the set
func NewInt64Set(values ...int64) *Int64Set {
	set := &Int64Set{items: make(map[int64]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Int64Set) Add(items ...int64) {
	for _, item := range items {
		set.items[item] = itemInt64Exists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Int64Set) Remove(items ...int64) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Int64Set) Contains(items ...int64) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Int64Set) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Int64Set) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Int64Set) Clear() {
	set.items = make(map[int64]struct{})
}

// Values returns all items in the set.
func (set *Int64Set) Values() []int64 {
	values := make([]int64, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Int64Set) String() string {
	str := "Int64Set\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

// ToJSON outputs the JSON representation of the set.
func (set *Int64Set) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Int64Set) FromJSON(data []byte) error {
	var elements []int64
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

// UintSet holds elements in go's native map
type UintSet struct {
	items map[uint]struct{}
}

var itemUintExists = struct{}{}

// NewUintSet instantiates a new empty set and adds the passed values, if any, to the set
func NewUintSet(values ...uint) *UintSet {
	set := &UintSet{items: make(map[uint]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *UintSet) Add(items ...uint) {
	for _, item := range items {
		set.items[item] = itemUintExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *UintSet) Remove(items ...uint) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *UintSet) Contains(items ...uint) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *UintSet) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *UintSet) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *UintSet) Clear() {
	set.items = make(map[uint]struct{})
}

// Values returns all items in the set.
func (set *UintSet) Values() []uint {
	values := make([]uint, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *UintSet) String() string {
	str := "UintSet\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

// ToJSON outputs the JSON representation of the set.
func (set *UintSet) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *UintSet) FromJSON(data []byte) error {
	var elements []uint
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

// Uint32Set holds elements in go's native map
type Uint32Set struct {
	items map[uint32]struct{}
}

var itemUint32Exists = struct{}{}

// NewUint32Set instantiates a new empty set and adds the passed values, if any, to the set
func NewUint32Set(values ...uint32) *Uint32Set {
	set := &Uint32Set{items: make(map[uint32]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Uint32Set) Add(items ...uint32) {
	for _, item := range items {
		set.items[item] = itemUint32Exists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Uint32Set) Remove(items ...uint32) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Uint32Set) Contains(items ...uint32) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Uint32Set) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Uint32Set) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Uint32Set) Clear() {
	set.items = make(map[uint32]struct{})
}

// Values returns all items in the set.
func (set *Uint32Set) Values() []uint32 {
	values := make([]uint32, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Uint32Set) String() string {
	str := "Uint32Set\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

// ToJSON outputs the JSON representation of the set.
func (set *Uint32Set) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Uint32Set) FromJSON(data []byte) error {
	var elements []uint32
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

// Uint64Set holds elements in go's native map
type Uint64Set struct {
	items map[uint64]struct{}
}

var itemUint64Exists = struct{}{}

// NewUint64Set instantiates a new empty set and adds the passed values, if any, to the set
func NewUint64Set(values ...uint64) *Uint64Set {
	set := &Uint64Set{items: make(map[uint64]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Uint64Set) Add(items ...uint64) {
	for _, item := range items {
		set.items[item] = itemUint64Exists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Uint64Set) Remove(items ...uint64) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Uint64Set) Contains(items ...uint64) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Uint64Set) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Uint64Set) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Uint64Set) Clear() {
	set.items = make(map[uint64]struct{})
}

// Values returns all items in the set.
func (set *Uint64Set) Values() []uint64 {
	values := make([]uint64, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Uint64Set) String() string {
	str := "Uint64Set\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

// ToJSON outputs the JSON representation of the set.
func (set *Uint64Set) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Uint64Set) FromJSON(data []byte) error {
	var elements []uint64
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

// Float32Set holds elements in go's native map
type Float32Set struct {
	items map[float32]struct{}
}

var itemFloat32Exists = struct{}{}

// NewFloat32Set instantiates a new empty set and adds the passed values, if any, to the set
func NewFloat32Set(values ...float32) *Float32Set {
	set := &Float32Set{items: make(map[float32]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Float32Set) Add(items ...float32) {
	for _, item := range items {
		set.items[item] = itemFloat32Exists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Float32Set) Remove(items ...float32) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Float32Set) Contains(items ...float32) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Float32Set) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Float32Set) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Float32Set) Clear() {
	set.items = make(map[float32]struct{})
}

// Values returns all items in the set.
func (set *Float32Set) Values() []float32 {
	values := make([]float32, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Float32Set) String() string {
	str := "Float32Set\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

// ToJSON outputs the JSON representation of the set.
func (set *Float32Set) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Float32Set) FromJSON(data []byte) error {
	var elements []float32
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29

// Float64Set holds elements in go's native map
type Float64Set struct {
	items map[float64]struct{}
}

var itemFloat64Exists = struct{}{}

// NewFloat64Set instantiates a new empty set and adds the passed values, if any, to the set
func NewFloat64Set(values ...float64) *Float64Set {
	set := &Float64Set{items: make(map[float64]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Float64Set) Add(items ...float64) {
	for _, item := range items {
		set.items[item] = itemFloat64Exists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Float64Set) Remove(items ...float64) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Float64Set) Contains(items ...float64) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Float64Set) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Float64Set) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Float64Set) Clear() {
	set.items = make(map[float64]struct{})
}

// Values returns all items in the set.
func (set *Float64Set) Values() []float64 {
	values := make([]float64, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Float64Set) String() string {
	str := "Float64Set\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

// ToJSON outputs the JSON representation of the set.
func (set *Float64Set) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Float64Set) FromJSON(data []byte) error {
	var elements []float64
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}
