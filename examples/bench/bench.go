// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"github.com/kelindar/column"
	"github.com/kelindar/column/fixtures"
	"github.com/kelindar/xxrand"
)

var (
	classes = []string{"fighter", "mage", "rogue"}
	races   = []string{"human", "elf", "dwarf", "orc"}
)

func main() {
	amount := 1000000
	players := column.NewCollection(column.Options{
		Capacity: amount,
	})
	createCollection(players, amount)

	// This runs point query benchmarks
	runBenchmark("Point Reads/Writes", func(writeTxn bool) (reads int, writes int) {

		// To avoid task granuarity problem, load up a bit more work on each
		// of the goroutines, a few hundred reads should be enough to amortize
		// the cost of scheduling goroutines, so we can actually test our code.
		for i := 0; i < 1000; i++ {
			offset := xxrand.Uint32n(uint32(amount - 1))
			if writeTxn {
				players.QueryAt(offset, func(r column.Row) error {
					r.SetFloat64("balance", 0)
					return nil
				})
				writes++
			} else {
				players.QueryAt(offset, func(r column.Row) error {
					_, _ = r.Float64("balance")
					return nil
				})
				reads++
			}
		}
		return
	})
}

// runBenchmark runs a benchmark
func runBenchmark(name string, fn func(bool) (int, int)) { _ = "STUB: not implemented"; return }

// Iterate over various concurrency levels

// createCollection loads a collection of players
func createCollection(out *column.Collection, amount int) *column.Collection {
	_ = "STUB: not implemented"
	return nil
}

// Load the data in

// insertPlayers inserts players
func insertPlayers(dst *column.Collection, data []fixtures.Player) error {
	_ = "STUB: not implemented"
	return nil
}
