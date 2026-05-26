// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package commit

import (
	"github.com/kelindar/bitmap"
)

const (
	size0    = 0      // 0 byte in size
	size2    = 1 << 4 // 2 bytes in size
	size4    = 2 << 4 // 4 bytes in size
	size8    = 3 << 4 // 8 bytes in size
	isNext   = 1 << 7 // is immediate next
	isString = 1 << 6 // is variable-size string
)

// --------------------------- Operation Type ----------------------------

// OpType represents a type of an operation.
type OpType uint8

// Various update operations supported.
const (
	Delete   OpType = 0 // Delete deletes an entire row or a set of rows
	Insert   OpType = 1 // Insert inserts a new row or a set of rows
	PutFalse OpType = 0 // PutFalse is a combination of Put+False for boolean values
	PutTrue  OpType = 2 // PutTrue is a combination of Put+True for boolean values
	Put      OpType = 2 // Put stores a value regardless of a previous value
	Merge    OpType = 3 // Applies a merge function
	Skip     OpType = 4 // Skips the value
)

// String returns a string representation
func (o OpType) String() string { _ = "STUB: not implemented"; return "" }

// --------------------------- Delta log ----------------------------

// Buffer represents a buffer of delta operations.
type Buffer struct {
	last   int32    // The last offset written
	chunk  Chunk    // The current chunk
	buffer []byte   // The destination buffer
	chunks []header // The offsets of chunks
	_      [8]byte  // padding
	Column string   // The column for the queue
}

// header represents a chunk metadata header.
type header struct {
	Chunk Chunk  // The chunk number
	Start uint32 // The offset at which the chunk starts in the buffer
	Value uint32 // The previous offset value for delta
}

// NewBuffer creates a new queue to store individual operations.
func NewBuffer(capacity int) *Buffer { _ = "STUB: not implemented"; return nil }

// Clone clones the buffer
func (b *Buffer) Clone() *Buffer { _ = "STUB: not implemented"; return nil }

// Reset resets the queue so it can be reused.
func (b *Buffer) Reset(column string) { _ = "STUB: not implemented"; return }

// IsEmpty returns whether the buffer is empty or not.
func (b *Buffer) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// Range iterates over the chunks present in the buffer
func (b *Buffer) RangeChunks(fn func(chunk Chunk)) { _ = "STUB: not implemented"; return }

// PutAny appends a supported value onto the buffer.
func (b *Buffer) PutAny(op OpType, idx uint32, value any) error {
	_ = "STUB: not implemented"
	return nil
}

// --------------------------- Numbers ----------------------------

// PutUint64 appends an uint64 value.
func (b *Buffer) PutUint64(op OpType, idx uint32, value uint64) { _ = "STUB: not implemented"; return }

// PutUint32 appends an uint32 value.
func (b *Buffer) PutUint32(op OpType, idx uint32, value uint32) { _ = "STUB: not implemented"; return }

// PutUint16 appends an uint16 value.
func (b *Buffer) PutUint16(op OpType, idx uint32, value uint16) { _ = "STUB: not implemented"; return }

// PutUint appends a uint64 value.
func (b *Buffer) PutUint(op OpType, idx uint32, value uint) { _ = "STUB: not implemented"; return }

// PutInt64 appends an int64 value.
func (b *Buffer) PutInt64(op OpType, idx uint32, value int64) { _ = "STUB: not implemented"; return }

// PutInt32 appends an int32 value.
func (b *Buffer) PutInt32(op OpType, idx uint32, value int32) { _ = "STUB: not implemented"; return }

// PutInt16 appends an int16 value.
func (b *Buffer) PutInt16(op OpType, idx uint32, value int16) { _ = "STUB: not implemented"; return }

// PutInt appends a int64 value.
func (b *Buffer) PutInt(op OpType, idx uint32, value int) { _ = "STUB: not implemented"; return }

// PutFloat64 appends a float64 value.
func (b *Buffer) PutFloat64(op OpType, idx uint32, value float64) {
	_ = "STUB: not implemented"
	return
}

// PutFloat32 appends an int32 value.
func (b *Buffer) PutFloat32(op OpType, idx uint32, value float32) {
	_ = "STUB: not implemented"
	return
}

// PutNumber appends a float64 value.
func (b *Buffer) PutNumber(op OpType, idx uint32, value float64) { _ = "STUB: not implemented"; return }

// --------------------------- Others ----------------------------

// PutOperation appends an operation type without a value.
func (b *Buffer) PutOperation(op OpType, idx uint32) { _ = "STUB: not implemented"; return }

// PutBool appends a boolean value.
func (b *Buffer) PutBool(idx uint32, value bool) {
	_ = "STUB: not implemented"

	// let the compiler do its magic: https://github.com/golang/go/issues/6011
	return
}

// PutBytes appends a binary value.
func (b *Buffer) PutBytes(op OpType, idx uint32, value []byte) { _ = "STUB: not implemented"; return }

// max 65K slices

// Write the the data itself and the offset

// PutString appends a string value.
func (b *Buffer) PutString(op OpType, idx uint32, value string) { _ = "STUB: not implemented"; return }

// PutBitmap iterates over the bitmap values and appends an operation for each bit set to one
func (b *Buffer) PutBitmap(op OpType, chunk Chunk, value bitmap.Bitmap) {
	_ = "STUB: not implemented"
	return
}

// writeUint64 appends a uint64 value.
func (b *Buffer) writeUint64(op OpType, idx uint32, value uint64) {
	_ = "STUB: not implemented"
	return
}

// writeUint32 appends a uint32 value.
func (b *Buffer) writeUint32(op OpType, idx uint32, value uint32) {
	_ = "STUB: not implemented"
	return
}

// writeUint16 appends a uint16 value.
func (b *Buffer) writeUint16(op OpType, idx uint32, value uint16) {
	_ = "STUB: not implemented"
	return
}

// writeOffset writes the offset at the current head.
func (b *Buffer) writeOffset(delta uint32) { _ = "STUB: not implemented"; return }

// writeChunk writes a chunk if changed and returns the delta
func (b *Buffer) writeChunk(idx uint32) int32 { _ = "STUB: not implemented"; return 0 }
