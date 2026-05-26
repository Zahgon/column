// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package column

import (
	"encoding"
	"sync"

	"github.com/kelindar/column/commit"
)

type recordType interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

// --------------------------- Record ----------------------------

// columnRecord represents a typed column that is persisted using binary marshaler
type columnRecord struct {
	columnString
	pool *sync.Pool
}

// ForRecord creates a new column that contains a type marshaled into/from binary. It requires
// a constructor for the type as well as optional merge function. If merge function is
// set to nil, "overwrite" strategy will be used.
func ForRecord[T recordType](new func() T, opts ...func(*option[T])) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// Merge function that decodes, merges and re-encodes records into their
// respective binary representation.

// Unmarshal the existing value

// Apply the user-defined merging strategy and marshal it back

// Value returns the value at the given index
// TODO: should probably get rid of this and use an `rdRecord` instead
func (c *columnRecord) Value(idx uint32) (out any, has bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// --------------------------- Writer ----------------------------

// rwRecord represents read-write accessor for primary keys.
type rwRecord struct {
	rdRecord
	writer *commit.Buffer
}

// Set sets the value at the current transaction index
func (s rwRecord) Set(value encoding.BinaryMarshaler) error { _ = "STUB: not implemented"; return nil }

// Merge atomically merges a delta to the value at the current transaction cursor
func (s rwRecord) Merge(delta encoding.BinaryMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

// write writes the operation
func (s rwRecord) write(op commit.OpType, encodeDelta func() ([]byte, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// As creates a read-write accessor for a specific record type.
func (txn *Txn) Record(columnName string) rwRecord {
	_ = "STUB: not implemented"
	return *new(rwRecord)
}

// --------------------------- Reader ----------------------------

// rdRecord represents a read-only accessor for records
type rdRecord reader[*columnRecord]

// Get loads the value at the current transaction index
func (s rdRecord) Get() (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

// Unmarshal loads the value at the current transaction index using a
// specified function to decode the value.
func (s rdRecord) Unmarshal(decode func(data []byte) error) bool {
	_ = "STUB: not implemented"
	return false
}

// readRecordOf creates a read-only accessor for readers
func readRecordOf(txn *Txn, columnName string) rdRecord {
	_ = "STUB: not implemented"
	return *new(rdRecord)
}

// --------------------------- Convert ----------------------------

// b2s converts byte slice to a string without allocating.
func b2s(b *[]byte) string { _ = "STUB: not implemented"; return "" }

// s2b converts a string to a byte slice without allocating.
func s2b(v string) (b []byte) { _ = "STUB: not implemented"; return nil }
