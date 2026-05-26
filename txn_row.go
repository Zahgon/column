// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package column

import (
	"encoding"
)

// Row represents a cursor at a particular row offest in the transaction.
type Row struct {
	txn *Txn
}

// Index returns the index of the row
func (r Row) Index() uint32 { _ = "STUB: not implemented"; return 0 }

// --------------------------- Numbers ----------------------------

// Int loads a int value at a particular column
func (r Row) Int(columnName string) (v int, ok bool) { _ = "STUB: not implemented"; return 0, false }

// SetInt stores a int value at a particular column
func (r Row) SetInt(columnName string, value int) { _ = "STUB: not implemented"; return }

// MergeInt atomically merges a delta into int value at a particular column
func (r Row) MergeInt(columnName string, value int) { _ = "STUB: not implemented"; return }

// Int16 loads a int16 value at a particular column
func (r Row) Int16(columnName string) (v int16, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetInt16 stores a int16 value at a particular column
func (r Row) SetInt16(columnName string, value int16) { _ = "STUB: not implemented"; return }

// MergeInt16 atomically merges a delta into int16 value at a particular column
func (r Row) MergeInt16(columnName string, value int16) { _ = "STUB: not implemented"; return }

// Int32 loads a int32 value at a particular column
func (r Row) Int32(columnName string) (v int32, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetInt32 stores a int32 value at a particular column
func (r Row) SetInt32(columnName string, value int32) { _ = "STUB: not implemented"; return }

// MergeInt32 atomically merges a delta into int32 value at a particular column
func (r Row) MergeInt32(columnName string, value int32) { _ = "STUB: not implemented"; return }

// Int64 loads a int64 value at a particular column
func (r Row) Int64(columnName string) (v int64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetInt64 stores a int64 value at a particular column
func (r Row) SetInt64(columnName string, value int64) { _ = "STUB: not implemented"; return }

// MergeInt64 atomically merges a delta into int64 value at a particular column
func (r Row) MergeInt64(columnName string, value int64) { _ = "STUB: not implemented"; return }

// Uint loads a uint value at a particular column
func (r Row) Uint(columnName string) (v uint, ok bool) { _ = "STUB: not implemented"; return 0, false }

// SetUint stores a uint value at a particular column
func (r Row) SetUint(columnName string, value uint) { _ = "STUB: not implemented"; return }

// MergeUint atomically merges a delta into uint value at a particular column
func (r Row) MergeUint(columnName string, value uint) { _ = "STUB: not implemented"; return }

// Uint16 loads a uint16 value at a particular column
func (r Row) Uint16(columnName string) (v uint16, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetUint16 stores a uint16 value at a particular column
func (r Row) SetUint16(columnName string, value uint16) { _ = "STUB: not implemented"; return }

// MergeUint16 atomically merges a delta into uint16 value at a particular column
func (r Row) MergeUint16(columnName string, value uint16) { _ = "STUB: not implemented"; return }

// Uint32 loads a uint32 value at a particular column
func (r Row) Uint32(columnName string) (v uint32, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetUint32 stores a uint32 value at a particular column
func (r Row) SetUint32(columnName string, value uint32) { _ = "STUB: not implemented"; return }

// MergeUint32 atomically merges a delta into uint32 value at a particular column
func (r Row) MergeUint32(columnName string, value uint32) { _ = "STUB: not implemented"; return }

// Uint64 loads a uint64 value at a particular column
func (r Row) Uint64(columnName string) (v uint64, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetUint64 stores a uint64 value at a particular column
func (r Row) SetUint64(columnName string, value uint64) { _ = "STUB: not implemented"; return }

// MergeUint64 atomically merges a delta into uint64 value at a particular column
func (r Row) MergeUint64(columnName string, value uint64) { _ = "STUB: not implemented"; return }

// Float32 loads a float32 value at a particular column
func (r Row) Float32(columnName string) (v float32, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// SetFloat32 stores a float32 value at a particular column
func (r Row) SetFloat32(columnName string, value float32) { _ = "STUB: not implemented"; return }

// MergeFloat32 atomically merges a delta into float32 value at a particular column
func (r Row) MergeFloat32(columnName string, value float32) { _ = "STUB: not implemented"; return }

// Float64 loads a float64 value at a particular column
func (r Row) Float64(columnName string) (float64, bool) { _ = "STUB: not implemented"; return 0, false }

// SetFloat64 stores a float64 value at a particular column
func (r Row) SetFloat64(columnName string, value float64) { _ = "STUB: not implemented"; return }

// MergeFloat64 atomically merges a delta into float64 value at a particular column
func (r Row) MergeFloat64(columnName string, value float64) { _ = "STUB: not implemented"; return }

// --------------------------- Strings ----------------------------

// Key loads a primary key value at a particular column
func (r Row) Key() (v string, ok bool) { _ = "STUB: not implemented"; return "", false }

// SetKey stores a primary key value at a particular column
func (r Row) SetKey(key string) { _ = "STUB: not implemented"; return }

// String loads a string value at a particular column
func (r Row) String(columnName string) (v string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

// SetString stores a string value at a particular column
func (r Row) SetString(columnName string, value string) { _ = "STUB: not implemented"; return }

// MergeString merges a string value at a particular column
func (r Row) MergeString(columnName string, value string) { _ = "STUB: not implemented"; return }

// Enum loads a string value at a particular column
func (r Row) Enum(columnName string) (v string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

// SetEnum stores a string value at a particular column
func (r Row) SetEnum(columnName string, value string) { _ = "STUB: not implemented"; return }

// --------------------------- Records ----------------------------

// Record loads a record value at a particular column
func (r Row) Record(columnName string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// SetRecord stores a record value at a particular column
func (r Row) SetRecord(columnName string, value encoding.BinaryMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

// MergeRecord merges a record value at a particular column
func (r Row) MergeRecord(columnName string, delta encoding.BinaryMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

// --------------------------- Map ----------------------------

// SetMany stores a set of columns for a given map
func (r Row) SetMany(value map[string]any) error { _ = "STUB: not implemented"; return nil }

// --------------------------- Others ----------------------------

// Bool loads a bool value at a particular column
func (r Row) Bool(columnName string) bool { _ = "STUB: not implemented"; return false }

// SetBool stores a bool value at a particular column
func (r Row) SetBool(columnName string, value bool) { _ = "STUB: not implemented"; return }

// Any loads a bool value at a particular column
func (r Row) Any(columnName string) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

// SetAny stores a bool value at a particular column
func (r Row) SetAny(columnName string, value interface{}) { _ = "STUB: not implemented"; return }
