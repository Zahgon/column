// This code was generated, DO NOT EDIT.
// Any changes will be lost if this file is regenerated.

package column

import (
	"github.com/kelindar/column/commit"
)

// --------------------------- Int ----------------------------

// makeInts creates a new vector for ints
func makeInts(opts ...func(*option[int])) Column { _ = "STUB: not implemented"; return *new(Column) }

// rwInt represents a read-write cursor for int
type rwInt struct {
	rdNumber[int]
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwInt) Set(value int) { _ = "STUB: not implemented"; return }

// Merge atomically merges a delta to the value at the current transaction cursor
func (s rwInt) Merge(delta int) { _ = "STUB: not implemented"; return }

// Int returns a read-write accessor for int column
func (txn *Txn) Int(columnName string) rwInt { _ = "STUB: not implemented"; return *new(rwInt) }

// --------------------------- Int16 ----------------------------

// makeInt16s creates a new vector for int16s
func makeInt16s(opts ...func(*option[int16])) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// rwInt16 represents a read-write cursor for int16
type rwInt16 struct {
	rdNumber[int16]
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwInt16) Set(value int16) { _ = "STUB: not implemented"; return }

// Merge atomically merges a delta to the value at the current transaction cursor
func (s rwInt16) Merge(delta int16) { _ = "STUB: not implemented"; return }

// Int16 returns a read-write accessor for int16 column
func (txn *Txn) Int16(columnName string) rwInt16 { _ = "STUB: not implemented"; return *new(rwInt16) }

// --------------------------- Int32 ----------------------------

// makeInt32s creates a new vector for int32s
func makeInt32s(opts ...func(*option[int32])) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// rwInt32 represents a read-write cursor for int32
type rwInt32 struct {
	rdNumber[int32]
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwInt32) Set(value int32) { _ = "STUB: not implemented"; return }

// Merge atomically merges a delta to the value at the current transaction cursor
func (s rwInt32) Merge(delta int32) { _ = "STUB: not implemented"; return }

// Int32 returns a read-write accessor for int32 column
func (txn *Txn) Int32(columnName string) rwInt32 { _ = "STUB: not implemented"; return *new(rwInt32) }

// --------------------------- Int64 ----------------------------

// makeInt64s creates a new vector for int64s
func makeInt64s(opts ...func(*option[int64])) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// rwInt64 represents a read-write cursor for int64
type rwInt64 struct {
	rdNumber[int64]
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwInt64) Set(value int64) { _ = "STUB: not implemented"; return }

// Merge atomically merges a delta to the value at the current transaction cursor
func (s rwInt64) Merge(delta int64) { _ = "STUB: not implemented"; return }

// Int64 returns a read-write accessor for int64 column
func (txn *Txn) Int64(columnName string) rwInt64 { _ = "STUB: not implemented"; return *new(rwInt64) }

// --------------------------- Uint ----------------------------

// makeUints creates a new vector for uints
func makeUints(opts ...func(*option[uint])) Column { _ = "STUB: not implemented"; return *new(Column) }

// rwUint represents a read-write cursor for uint
type rwUint struct {
	rdNumber[uint]
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwUint) Set(value uint) { _ = "STUB: not implemented"; return }

// Merge atomically merges a delta to the value at the current transaction cursor
func (s rwUint) Merge(delta uint) { _ = "STUB: not implemented"; return }

// Uint returns a read-write accessor for uint column
func (txn *Txn) Uint(columnName string) rwUint { _ = "STUB: not implemented"; return *new(rwUint) }

// --------------------------- Uint16 ----------------------------

// makeUint16s creates a new vector for uint16s
func makeUint16s(opts ...func(*option[uint16])) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// rwUint16 represents a read-write cursor for uint16
type rwUint16 struct {
	rdNumber[uint16]
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwUint16) Set(value uint16) { _ = "STUB: not implemented"; return }

// Merge atomically merges a delta to the value at the current transaction cursor
func (s rwUint16) Merge(delta uint16) { _ = "STUB: not implemented"; return }

// Uint16 returns a read-write accessor for uint16 column
func (txn *Txn) Uint16(columnName string) rwUint16 {
	_ = "STUB: not implemented"
	return *new(rwUint16)
}

// --------------------------- Uint32 ----------------------------

// makeUint32s creates a new vector for uint32s
func makeUint32s(opts ...func(*option[uint32])) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// rwUint32 represents a read-write cursor for uint32
type rwUint32 struct {
	rdNumber[uint32]
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwUint32) Set(value uint32) { _ = "STUB: not implemented"; return }

// Merge atomically merges a delta to the value at the current transaction cursor
func (s rwUint32) Merge(delta uint32) { _ = "STUB: not implemented"; return }

// Uint32 returns a read-write accessor for uint32 column
func (txn *Txn) Uint32(columnName string) rwUint32 {
	_ = "STUB: not implemented"
	return *new(rwUint32)
}

// --------------------------- Uint64 ----------------------------

// makeUint64s creates a new vector for uint64s
func makeUint64s(opts ...func(*option[uint64])) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// rwUint64 represents a read-write cursor for uint64
type rwUint64 struct {
	rdNumber[uint64]
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwUint64) Set(value uint64) { _ = "STUB: not implemented"; return }

// Merge atomically merges a delta to the value at the current transaction cursor
func (s rwUint64) Merge(delta uint64) { _ = "STUB: not implemented"; return }

// Uint64 returns a read-write accessor for uint64 column
func (txn *Txn) Uint64(columnName string) rwUint64 {
	_ = "STUB: not implemented"
	return *new(rwUint64)
}

// --------------------------- Float32 ----------------------------

// makeFloat32s creates a new vector for float32s
func makeFloat32s(opts ...func(*option[float32])) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// rwFloat32 represents a read-write cursor for float32
type rwFloat32 struct {
	rdNumber[float32]
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwFloat32) Set(value float32) { _ = "STUB: not implemented"; return }

// Merge atomically merges a delta to the value at the current transaction cursor
func (s rwFloat32) Merge(delta float32) { _ = "STUB: not implemented"; return }

// Float32 returns a read-write accessor for float32 column
func (txn *Txn) Float32(columnName string) rwFloat32 {
	_ = "STUB: not implemented"
	return *new(rwFloat32)
}

// --------------------------- Float64 ----------------------------

// makeFloat64s creates a new vector for float64s
func makeFloat64s(opts ...func(*option[float64])) Column {
	_ = "STUB: not implemented"
	return *new(Column)
}

// rwFloat64 represents a read-write cursor for float64
type rwFloat64 struct {
	rdNumber[float64]
	writer *commit.Buffer
}

// Set sets the value at the current transaction cursor
func (s rwFloat64) Set(value float64) { _ = "STUB: not implemented"; return }

// Merge atomically merges a delta to the value at the current transaction cursor
func (s rwFloat64) Merge(delta float64) { _ = "STUB: not implemented"; return }

// Float64 returns a read-write accessor for float64 column
func (txn *Txn) Float64(columnName string) rwFloat64 {
	_ = "STUB: not implemented"
	return *new(rwFloat64)
}
