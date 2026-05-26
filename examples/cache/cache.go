// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"github.com/kelindar/column"
)

// Cache represents a key-value store
type Cache struct {
	store *column.Collection
}

// New creates a new key-value cache
func New() *Cache { _ = "STUB: not implemented"; return nil }

// Get attempts to retrieve a value for a key
func (c *Cache) Get(key string) (value string, found bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Set updates or inserts a new value
func (c *Cache) Set(key, value string) { _ = "STUB: not implemented"; return }
