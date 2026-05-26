// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package column

import (
	"reflect"
	"sync"

	"github.com/kelindar/bitmap"
	"github.com/kelindar/column/commit"
)

// columnType represents a type of a column.
type columnType uint8

const (
	typeGeneric = columnType(0)      // Generic column, every column should support this
	typeNumeric = columnType(1 << 0) // Numeric column supporting float64, int64 or uint64
	typeTextual = columnType(1 << 1) // Textual column supporting strings
)

// typeOf resolves all supported types of the column
func typeOf(column Column) (typ columnType) { _ = "STUB: not implemented"; return *new(columnType) }

// --------------------------- Contracts ----------------------------

// Column represents a column implementation
type Column interface {
	Grow(idx uint32)
	Apply(commit.Chunk, *commit.Reader)
	Value(idx uint32) (interface{}, bool)
	Contains(idx uint32) bool
	Index(commit.Chunk) bitmap.Bitmap
	Snapshot(chunk commit.Chunk, dst *commit.Buffer)
}

// Numeric represents a column that stores numbers.
type Numeric interface {
	Column
	LoadFloat64(uint32) (float64, bool)
	LoadUint64(uint32) (uint64, bool)
	LoadInt64(uint32) (int64, bool)
	FilterFloat64(commit.Chunk, bitmap.Bitmap, func(v float64) bool)
	FilterUint64(commit.Chunk, bitmap.Bitmap, func(v uint64) bool)
	FilterInt64(commit.Chunk, bitmap.Bitmap, func(v int64) bool)
}

// Textual represents a column that stores strings.
type Textual interface {
	Column
	LoadString(uint32) (string, bool)
	FilterString(commit.Chunk, bitmap.Bitmap, func(v string) bool)
}

// --------------------------- Constructors ----------------------------

// Various column constructor functions for a specific types.
var (
	ForString  = makeStrings
	ForFloat32 = makeFloat32s
	ForFloat64 = makeFloat64s
	ForInt     = makeInts
	ForInt16   = makeInt16s
	ForInt32   = makeInt32s
	ForInt64   = makeInt64s
	ForUint    = makeUints
	ForUint16  = makeUint16s
	ForUint32  = makeUint32s
	ForUint64  = makeUint64s
	ForBool    = makeBools
	ForEnum    = makeEnum
	ForKey     = makeKey
)

// ForKind creates a new column instance for a specified reflect.Kind
func ForKind(kind reflect.Kind) (Column, error) {
	_ = "STUB: not implemented"
	return *new(Column), nil
}

// --------------------------- Generic Options ----------------------------

// option represents options for variouos columns.
type option[T any] struct {
	Merge func(value, delta T) T
}

// configure applies options
func configure[T any](opts []func(*option[T]), dst option[T]) option[T] {
	_ = "STUB: not implemented"
	return nil
}

// WithMerge sets an optional merge function that allows you to merge a delta value to
// an existing value, atomically. The operation is performed transactionally.
func WithMerge[T any](fn func(value, delta T) T) func(*option[T]) {
	_ = "STUB: not implemented"
	return nil
}

// --------------------------- Column ----------------------------

// column represents a column wrapper that synchronizes operations
type column struct {
	Column
	lock sync.RWMutex // The lock to protect the entire column
	kind columnType   // The type of the colum
	name string       // The name of the column
}

// columnFor creates a synchronized column for a column implementation
func columnFor(name string, v Column) *column { _ = "STUB: not implemented"; return nil }

// IsIndex returns whether the column is an index
func (c *column) IsIndex() bool { _ = "STUB: not implemented"; return false }

// IsNumeric checks whether a column type supports certain numerical operations.
func (c *column) IsNumeric() bool { _ = "STUB: not implemented"; return false }

// IsTextual checks whether a column type supports certain string operations.
func (c *column) IsTextual() bool { _ = "STUB: not implemented"; return false }

// Grow grows the size of the column
func (c *column) Grow(idx uint32) { _ = "STUB: not implemented"; return }

// Apply performs a series of operations on a column.
func (c *column) Apply(chunk commit.Chunk, r *commit.Reader) { _ = "STUB: not implemented"; return }

// Index loads the appropriate column index for a given chunk
func (c *column) Index(chunk commit.Chunk) bitmap.Bitmap {
	_ = "STUB: not implemented"
	return *new(bitmap.Bitmap)
}

// Snapshot takes a snapshot of a column, skipping indexes
func (c *column) Snapshot(chunk commit.Chunk, buffer *commit.Buffer) bool {
	_ = "STUB: not implemented"
	return false
}

// Value retrieves a value at a specified index
func (c *column) Value(idx uint32) (v interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// --------------------------- Accessor  ----------------------------

// Reader represents a generic reader
type reader[T any] struct {
	cursor *uint32
	reader T
}

// readerFor creates a read-only accessor
func readerFor[T any](txn *Txn, columnName string) reader[T] { _ = "STUB: not implemented"; return nil }

// --------------------------- Any Writer ----------------------------

// rwAny represents read-write accessor for any column type
type rwAny struct {
	rdAny
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwAny) Set(value any) error { _ = "STUB: not implemented"; return nil }

// --------------------------- Any Reader ----------------------------

// rdAny represents a read-only accessor for any value
type rdAny reader[Column]

// Get loads the value at the current transaction cursor
func (s rdAny) Get() (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

// readAnyOf creates a new any reader
func readAnyOf(txn *Txn, columnName string) rdAny { _ = "STUB: not implemented"; return *new(rdAny) }

// Any returns a column accessor
func (txn *Txn) Any(columnName string) rwAny { _ = "STUB: not implemented"; return *new(rwAny) }

// --------------------------- segment list ----------------------------

// Chunks represents a chunked array storage
type chunks[T any] []struct {
	fill bitmap.Bitmap // The fill-list
	data []T           // The actual values
}

// chunkAt loads the fill and data list at a particular chunk
func (s chunks[T]) chunkAt(chunk commit.Chunk) (bitmap.Bitmap, []T) {
	_ = "STUB: not implemented"
	return *new(bitmap.Bitmap), nil
}

// Grow grows a segment list
func (s *chunks[T]) Grow(idx uint32) { _ = "STUB: not implemented"; return }

// Index returns the fill list for the segment
func (s chunks[T]) Index(chunk commit.Chunk) (fill bitmap.Bitmap) {
	_ = "STUB: not implemented"
	return *new(bitmap.Bitmap)
}
