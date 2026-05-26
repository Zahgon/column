// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package commit

import (
	"io"

	"github.com/kelindar/iostream"
)

// --------------------------- WriteTo ----------------------------

// WriteTo writes data to w until there's no more data to write or when an error occurs. The return
// value n is the number of bytes written. Any error encountered during the write is also returned.
func (b *Buffer) WriteTo(dst io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// --------------------------- ReadFrom ----------------------------

// ReadFrom reads data from r until EOF or error. The return value n is the number of
// bytes read. Any error except EOF encountered during the read is also returned.
func (b *Buffer) ReadFrom(src io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// readChunksFrom reads the list of chunks from the reader
func readChunksFrom(r *iostream.Reader) ([]header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// toBytes converts a string to a byte slice without allocating.
func toBytes(v string) (b []byte) { _ = "STUB: not implemented"; return nil }
