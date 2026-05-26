// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package column

import (
	"errors"
	"io"

	"github.com/kelindar/bitmap"
	"github.com/kelindar/column/commit"
)

var (
	errUnexpectedEOF = errors.New("column: unable to restore, unexpected EOF")
)

// --------------------------- Commit Replay ---------------------------

// Replay replays a commit on a collection, applying the changes.
func (c *Collection) Replay(change commit.Commit) error { _ = "STUB: not implemented"; return nil }

// --------------------------- Snapshotting ---------------------------

// Restore restores the collection from the underlying snapshot reader. This operation
// should be called before any of transactions, right after initialization.
func (c *Collection) Restore(snapshot io.Reader) error { _ = "STUB: not implemented"; return nil }

// Reconcile the pending commit log

// Snapshot writes a collection snapshot into the underlying writer.
func (c *Collection) Snapshot(dst io.Writer) error { _ = "STUB: not implemented"; return nil }

// Take a snapshot of the current state

// Close the recorder

// recorderOpen opens a recorder for commits while the snapshot is in progress
func (c *Collection) recorderOpen() (log *commit.Log, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// recorderClose closes the pending commit recorder and deletes the file
func (c *Collection) recorderClose() { _ = "STUB: not implemented"; return }

// isSnapshotting loads a currently used commit log for a pending snapshot
func (c *Collection) isSnapshotting() (*commit.Log, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// --------------------------- Collection Encoding ---------------------------

// writeState writes collection state into the specified writer.
func (c *Collection) writeState(dst io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Write the schema version

// Load the number of columns and the max index

// extra 'insert' column

// Write the number of columns

// Write each chunk

// Write the last written commit for this chunk

// Write the inserts column

// Snapshot each column and write the buffer

// Skip indexes

// readState reads a collection snapshotted state from the underlying reader. It
// returns the last commit IDs for each chunk.
func (c *Collection) readState(src io.Reader) (map[commit.Chunk]uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read the version and make sure it matches

// Read the number of columns

// Read each chunk

// Read the last written commit ID for the chunk

// chunks returns the number of chunks and columns
func (c *Collection) chunks() int { _ = "STUB: not implemented"; return 0 }

// readChunk acquires appropriate locks for a chunk and executes a read callback.
// This is used for snapshotting purposes only.
func (c *Collection) readChunk(chunk commit.Chunk, fn func(uint64, commit.Chunk, bitmap.Bitmap) error) error {
	_ = "STUB: not implemented"

	// Lock both the chunk and the fill list
	return nil
}
