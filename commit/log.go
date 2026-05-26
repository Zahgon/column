// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package commit

import (
	"io"
	"os"
	"sync"

	"github.com/kelindar/iostream"
)

// Logger represents a contract that a commit logger must implement
type Logger interface {
	Append(commit Commit) error
}

var _ Logger = new(Channel)
var _ Logger = new(Log)

// --------------------------- Channel ----------------------------

// Channel represents an impementation of a commit writer that simply sends each commit
// into the channel.
type Channel chan Commit

// Append clones the commit and writes it into the logger
func (w Channel) Append(commit Commit) error { _ = "STUB: not implemented"; return nil }

// --------------------------- Log ----------------------------

// Log represents a commit log that can be used to write the changes to the collection
// during a snapshot. It also supports reading a commit log back.
type Log struct {
	lock   sync.Mutex
	source io.Reader
	writer *iostream.Writer
	reader *iostream.Reader
}

// Open opens a commit log stream for both read and write.
func Open(source io.Reader) *Log { _ = "STUB: not implemented"; return nil }

// OpenFile opens a specified commit log file in a read/write mode. If
// the file does not exist, it will create it.
func OpenFile(filename string) (*Log, error) { _ = "STUB: not implemented"; return nil, nil }

// OpenTemp opens a temporary commit log file with read/write permissions
func OpenTemp() (*Log, error) { _ = "STUB: not implemented"; return nil, nil }

// openFile opens a file or returns the error provided
func openFile(file *os.File, err error) (*Log, error) { _ = "STUB: not implemented"; return nil, nil }

// Append writes the commit into the log destination
func (l *Log) Append(commit Commit) (err error) { _ = "STUB: not implemented"; return nil }

// Write the commit into the stream

// Range iterates over all the commits in the log and calls the provided
// callback function on each of them. If the callback returns an error, the
// iteration will stop.
func (l *Log) Range(fn func(Commit) error) error { _ = "STUB: not implemented"; return nil }

// Read the commit

// Name calls the corresponding Name() method on the underlying source
func (l *Log) Name() (name string) { _ = "STUB: not implemented"; return "" }

// Copy copies the contents of the log into the destination writer.
func (l *Log) Copy(dst io.Writer) error { _ = "STUB: not implemented"; return nil }

// Rewind to the beginning of the file, the underlying source must
// implement io.Seeker for this to work.

// Append the pending commits to the destination

// Close closes the source log file.
func (l *Log) Close() (err error) { _ = "STUB: not implemented"; return nil }
