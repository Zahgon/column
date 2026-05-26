// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package column

import (
	"sync"

	"github.com/kelindar/bitmap"
	"github.com/kelindar/column/commit"

	"github.com/tidwall/btree"
)

// --------------------------- Reader ---------------------------

// Reader represents a reader cursor for a specific row/column combination.
type Reader interface {
	IsUpsert() bool
	IsDelete() bool
	Index() uint32
	String() string
	Bytes() []byte
	Float() float64
	Int() int
	Uint() uint
	Bool() bool
}

// Assert reader implementations. Both our cursor and commit reader need to implement
// this so that we can feed it to the index transparently.
var _ Reader = new(commit.Reader)

// computed represents a computed column
type computed interface {
	Column() string
}

// --------------------------- Index ----------------------------

// columnIndex represents the index implementation
type columnIndex struct {
	fill bitmap.Bitmap     // The fill list for the column
	name string            // The name of the target column
	rule func(Reader) bool // The rule to apply when building the index
}

// newIndex creates a new bitmap index column.
func newIndex(indexName, columnName string, rule func(Reader) bool) *column {
	_ = "STUB: not implemented"
	return nil
}

// Grow grows the size of the column until we have enough to store
func (c *columnIndex) Grow(idx uint32) {
	_ = "STUB: not implemented"

	// Column returns the target name of the column on which this index should apply.
	return
}

func (c *columnIndex) Column() string {
	_ = "STUB: not implemented"

	// Apply applies a set of operations to the column.
	return ""
}

func (c *columnIndex) Apply(chunk commit.Chunk, r *commit.Reader) {
	_ = "STUB: not implemented"

	// Index can only be updated based on the final stored value, so we can only work
	// with put operations here. The trick is to update the final value after applying
	// on the actual column.
	return
}

// Value retrieves a value at a specified index.
func (c *columnIndex) Value(idx uint32) (v interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Contains checks whether the column has a value at a specified index.
func (c *columnIndex) Contains(idx uint32) bool { _ = "STUB: not implemented"; return false }

// Index returns the fill list for the column
func (c *columnIndex) Index(chunk commit.Chunk) bitmap.Bitmap {
	_ = "STUB: not implemented"
	return *new(bitmap.Bitmap)
}

// Snapshot writes the entire column into the specified destination buffer
func (c *columnIndex) Snapshot(chunk commit.Chunk, dst *commit.Buffer) {
	_ = "STUB: not implemented"
	return
}

// --------------------------- Trigger ----------------------------

// columnTrigger represents the trigger implementation
type columnTrigger struct {
	name string       // The name of the target column
	clbk func(Reader) // The trigger callback
}

// newTrigger creates a new trigger column.
func newTrigger(indexName, columnName string, callback func(r Reader)) *column {
	_ = "STUB: not implemented"
	return nil
}

// Column returns the target name of the column on which this index should apply.
func (c *columnTrigger) Column() string {
	_ = "STUB: not implemented"

	// Grow grows the size of the column until we have enough to store
	return ""
}

func (c *columnTrigger) Grow(idx uint32) {
	_ = "STUB: not implemented"
	// Noop

	// Apply applies a set of operations to the column.
	return
}

func (c *columnTrigger) Apply(chunk commit.Chunk, r *commit.Reader) {
	_ = "STUB: not implemented"
	return
}

// Value retrieves a value at a specified index.
func (c *columnTrigger) Value(idx uint32) (v any, ok bool) {
	_ = "STUB: not implemented"

	// Contains checks whether the column has a value at a specified index.
	return *new(any), false
}

func (c *columnTrigger) Contains(idx uint32) bool {
	_ = "STUB: not implemented"

	// Index returns the fill list for the column
	return false
}

func (c *columnTrigger) Index(chunk commit.Chunk) bitmap.Bitmap {
	_ = "STUB: not implemented"

	// Snapshot writes the entire column into the specified destination buffer
	return *new(bitmap.Bitmap)
}

func (c *columnTrigger) Snapshot(chunk commit.Chunk, dst *commit.Buffer) {
	_ = "STUB: not implemented"
	// Noop

	// ----------------------- Sorted Index --------------------------
	return
}

type sortIndexItem struct {
	Key   string
	Value uint32
}

// columnSortIndex implements a constantly sorted column via BTree
type columnSortIndex struct {
	btree    *btree.BTreeG[sortIndexItem] // 1 constantly sorted data structure
	backMap  map[uint32]string            // for constant key lookups
	backLock sync.Mutex                   // protect backMap access
	name     string                       // The name of the target column
}

// newSortIndex creates a new bitmap index column.
func newSortIndex(indexName, columnName string) *column { _ = "STUB: not implemented"; return nil }

// Grow grows the size of the column until we have enough to store
func (c *columnSortIndex) Grow(idx uint32) {
	_ = "STUB: not implemented"

	// Column returns the target name of the column on which this index should apply.
	return
}

func (c *columnSortIndex) Column() string {
	_ = "STUB: not implemented"

	// Apply applies a set of operations to the column.
	return ""
}

func (c *columnSortIndex) Apply(chunk commit.Chunk, r *commit.Reader) {
	_ = "STUB: not implemented"

	// Index can only be updated based on the final stored value, so we can only work
	// with put, merge, & delete operations here.
	return
}

// alloc required

// Value retrieves a value at a specified index.
func (c *columnSortIndex) Value(idx uint32) (v interface{}, ok bool) {
	_ = "STUB: not implemented"

	// Contains checks whether the column has a value at a specified index.
	return nil, false
}

func (c *columnSortIndex) Contains(idx uint32) bool {
	_ = "STUB: not implemented"

	// Index returns the fill list for the column
	return false
}

func (c *columnSortIndex) Index(chunk commit.Chunk) bitmap.Bitmap {
	_ = "STUB: not implemented"

	// Snapshot writes the entire column into the specified destination buffer
	return *new(bitmap.Bitmap)
}

func (c *columnSortIndex) Snapshot(chunk commit.Chunk, dst *commit.Buffer) {
	_ = "STUB: not implemented"
	// No-op
	return
}
