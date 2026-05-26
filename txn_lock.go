// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package column

import (
	"github.com/kelindar/bitmap"
	"github.com/kelindar/column/commit"
)

const (
	bitmapShift = chunkShift - 6
	bitmapSize  = 1 << bitmapShift
	chunkShift  = 14 // 16K
	chunkSize   = 1 << chunkShift
)

// initialize ensures that the transaction is pre-initialized with the snapshot
// of the owner's fill list.
func (txn *Txn) initialize() { _ = "STUB: not implemented"; return }

// --------------------------- Locked Seek ---------------------------

// QueryAt jumps at a particular offset in the collection, sets the cursor to the
// provided position and executes given callback fn.
func (txn *Txn) QueryAt(index uint32, f func(Row) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// --------------------------- Locked Range ---------------------------

// rangeRead iterates over index, chunk by chunk and ensures that each
// chunk is protected by an appropriate read lock.
func (txn *Txn) rangeRead(f func(chunk commit.Chunk, index bitmap.Bitmap)) {
	_ = "STUB: not implemented"
	return
}

// rangeReadPair iterates over the index and another bitmap, chunk by chunk and
// ensures that each chunk is protected by an appropriate read lock.
func (txn *Txn) rangeReadPair(column *column, f func(a, b bitmap.Bitmap)) {
	_ = "STUB: not implemented"
	return
}

// Iterate through all of the chunks and acquire appropriate shard locks.

// rangeWrite ranges over the dirty chunks and acquires exclusive latches along
// the way. This is used to commit a transaction.
func (txn *Txn) rangeWrite(fn func(commitID uint64, chunk commit.Chunk, fill bitmap.Bitmap)) {
	_ = "STUB: not implemented"
	return
}

// Compute the fill and set the last commit ID

// OK, since we have a shard lock

// Call the delegate
