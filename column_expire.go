// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package column

import (
	"context"
	"time"
)

// --------------------------- Expiration (Vacuum) ----------------------------

// vacuum cleans up the expired objects on a specified interval.
func (c *Collection) vacuum(ctx context.Context, interval time.Duration) {
	_ = "STUB: not implemented"
	return
}

// --------------------------- Expiration (Column) ----------------------------

// TTL returns a read-write accessor for the time-to-live column
func (txn *Txn) TTL() rwTTL { _ = "STUB: not implemented"; return *new(rwTTL) }

type rwTTL struct {
	rw rwInt64
}

// TTL returns the remaining time-to-live duration
func (s rwTTL) TTL() (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}

// ExpiresAt returns the expiration time
func (s rwTTL) ExpiresAt() (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

// Set sets the time-to-live value at the current transaction cursor
func (s rwTTL) Set(ttl time.Duration) { _ = "STUB: not implemented"; return }

// Extend extends time-to-live of the row current transaction cursor by a specified amount
func (s rwTTL) Extend(delta time.Duration) { _ = "STUB: not implemented"; return }

// readTTL converts expiration to a TTL
func readTTL(expireAt int64) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// writeTTL converts ttl to expireAt
func writeTTL(ttl time.Duration) int64 { _ = "STUB: not implemented"; return 0 }

// --------------------------- Expiration (Row) ----------------------------

// TTL retrieves the time left before the row will be cleaned up
func (r Row) TTL() (time.Duration, bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}

// SetTTL sets a time-to-live for a row and returns the expiration time
func (r Row) SetTTL(ttl time.Duration) (until time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// Otherwise, return zero time (never expires)
