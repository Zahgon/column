// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package column

import (
	"github.com/kelindar/bitmap"
	"github.com/kelindar/column/commit"
)

// columnBool represents a boolean column
type columnBool struct {
	data bitmap.Bitmap
}

// makeBools creates a new boolean column
func makeBools() Column { _ = "STUB: not implemented"; return *new(Column) }

// Grow grows the size of the column until we have enough to store
func (c *columnBool) Grow(idx uint32) {
	_ = "STUB: not implemented"

	// Apply applies a set of operations to the column.
	return
}

func (c *columnBool) Apply(chunk commit.Chunk, r *commit.Reader) { _ = "STUB: not implemented"; return }

// also "delete"

// Value retrieves a value at a specified index
func (c *columnBool) Value(idx uint32) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Contains checks whether the column has a value at a specified index.
func (c *columnBool) Contains(idx uint32) bool { _ = "STUB: not implemented"; return false }

// Index returns the fill list for the column
func (c *columnBool) Index(chunk commit.Chunk) bitmap.Bitmap {
	_ = "STUB: not implemented"
	return *new(bitmap.Bitmap)
}

// Snapshot writes the entire column into the specified destination buffer
func (c *columnBool) Snapshot(chunk commit.Chunk, dst *commit.Buffer) {
	_ = "STUB: not implemented"
	return
}

// --------------------------- Writer ----------------------------

// rwBool represents read-write accessor for boolean values
type rwBool struct {
	rdBool
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwBool) Set(value bool) { _ = "STUB: not implemented"; return }

// Bool returns a bool column accessor
func (txn *Txn) Bool(columnName string) rwBool { _ = "STUB: not implemented"; return *new(rwBool) }

// --------------------------- Reader ----------------------------

// rdBool represents a read-only accessor for boolean values
type rdBool reader[Column]

// Get loads the value at the current transaction cursor
func (s rdBool) Get() bool { _ = "STUB: not implemented"; return false }

// readBoolOf creates a new boolean reader
func readBoolOf(txn *Txn, columnName string) rdBool { _ = "STUB: not implemented"; return *new(rdBool) }
