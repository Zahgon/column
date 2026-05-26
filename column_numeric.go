// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package column

import (
	"github.com/kelindar/bitmap"
	"github.com/kelindar/column/commit"
	"github.com/kelindar/simd"
)

//go:generate go run ./codegen/main.go

// readNumber is a helper function for point reads
func readNumber[T simd.Number](txn *Txn, columnName string) (value T, found bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// --------------------------- Generic Column ----------------------------

// numericColumn represents a numeric column
type numericColumn[T simd.Number] struct {
	chunks[T]
	option[T]
	write func(*commit.Buffer, uint32, T)
	apply func(*commit.Reader, bitmap.Bitmap, []T, option[T])
}

// makeNumeric creates a new vector for simd.Numbers
func makeNumeric[T simd.Number](
	write func(*commit.Buffer, uint32, T),
	apply func(*commit.Reader, bitmap.Bitmap, []T, option[T]),
	opts []func(*option[T]),
) *numericColumn[T] {
	_ = "STUB: not implemented"
	return nil
}

// --------------------------- Accessors ----------------------------

// Contains checks whether the column has a value at a specified index.
func (c *numericColumn[T]) Contains(idx uint32) bool { _ = "STUB: not implemented"; return false }

// load retrieves a float64 value at a specified index
func (c *numericColumn[T]) load(idx uint32) (v T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Value retrieves a value at a specified index
func (c *numericColumn[T]) Value(idx uint32) (any, bool) {
	_ = "STUB: not implemented"
	return *

	// LoadFloat64 retrieves a float64 value at a specified index
	new(any), false
}

func (c *numericColumn[T]) LoadFloat64(idx uint32) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// LoadInt64 retrieves an int64 value at a specified index
func (c *numericColumn[T]) LoadInt64(idx uint32) (int64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// LoadUint64 retrieves an uint64 value at a specified index
func (c *numericColumn[T]) LoadUint64(idx uint32) (uint64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// --------------------------- Filtering ----------------------------

// filterNumbers filters down the values based on the specified predicate.
func filterNumbers[T, C simd.Number](column *numericColumn[T], chunk commit.Chunk, index bitmap.Bitmap, predicate func(C) bool) {
	_ = "STUB: not implemented"
	return
}

// FilterFloat64 filters down the values based on the specified predicate.
func (c *numericColumn[T]) FilterFloat64(chunk commit.Chunk, index bitmap.Bitmap, predicate func(float64) bool) {
	_ = "STUB: not implemented"
	return
}

// FilterInt64 filters down the values based on the specified predicate.
func (c *numericColumn[T]) FilterInt64(chunk commit.Chunk, index bitmap.Bitmap, predicate func(int64) bool) {
	_ = "STUB: not implemented"
	return
}

// FilterUint64 filters down the values based on the specified predicate.
func (c *numericColumn[T]) FilterUint64(chunk commit.Chunk, index bitmap.Bitmap, predicate func(uint64) bool) {
	_ = "STUB: not implemented"
	return
}

// --------------------------- Apply & Snapshot ----------------------------

// Apply applies a set of operations to the column.
func (c *numericColumn[T]) Apply(chunk commit.Chunk, r *commit.Reader) {
	_ = "STUB: not implemented"
	return
}

// Snapshot writes the entire column into the specified destination buffer
func (c *numericColumn[T]) Snapshot(chunk commit.Chunk, dst *commit.Buffer) {
	_ = "STUB: not implemented"
	return
}

// --------------------------- Reader/Writer ----------------------------

// rdNumber represents a read-only accessor for simd.Numbers
type rdNumber[T simd.Number] struct {
	reader *numericColumn[T]
	txn    *Txn
}

// Get loads the value at the current transaction cursor
func (s rdNumber[T]) Get() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// Sum computes a sum of the column values selected by this transaction
func (s rdNumber[T]) Sum() (sum T) { _ = "STUB: not implemented"; return *new(T) }

// Avg computes an arithmetic mean of the column values selected by this transaction
func (s rdNumber[T]) Avg() float64 { _ = "STUB: not implemented"; return 0 }

// Min finds the smallest value from the column values selected by this transaction
func (s rdNumber[T]) Min() (min T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

// Max finds the largest value from the column values selected by this transaction
func (s rdNumber[T]) Max() (max T, ok bool) { _ = "STUB: not implemented"; return *new(T), false }

// readNumberOf creates a new numeric reader
func readNumberOf[T simd.Number](txn *Txn, columnName string) rdNumber[T] {
	_ = "STUB: not implemented"
	return nil
}
