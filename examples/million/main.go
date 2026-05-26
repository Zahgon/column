// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"bytes"
	"fmt"

	"github.com/kelindar/column"
	"github.com/kelindar/column/fixtures"
)

func main() {
	amount, runs := 10000000, 20
	players := column.NewCollection(column.Options{
		Capacity: amount,
	})

	// insert the data first
	measure("insert", fmt.Sprintf("%v rows", amount), func() {
		createCollection(players, amount)
	}, 1)

	// snapshot the dataset
	measure("snapshot", fmt.Sprintf("%v rows", amount), func() {
		buffer := bytes.NewBuffer(nil)
		players.Snapshot(buffer)
	}, 10)

	// run a full scan
	measure("full scan", "age >= 30", func() {
		players.Query(func(txn *column.Txn) error {
			count := txn.WithFloat("age", func(v float64) bool {
				return v >= 30
			}).Count()
			fmt.Printf("-> result = %v\n", count)
			return nil
		})
	}, runs)

	// run a full scan
	measure("full scan", `class == "rogue"`, func() {
		players.Query(func(txn *column.Txn) error {
			count := txn.WithString("class", func(v string) bool {
				return v == "rogue"
			}).Count()
			fmt.Printf("-> result = %v\n", count)
			return nil
		})
	}, runs)

	// run a query over human mages
	measure("indexed query", "human mages", func() {
		players.Query(func(txn *column.Txn) error {
			fmt.Printf("-> result = %v\n", txn.With("human", "mage").Count())
			return nil
		})
	}, runs*1000)

	// run a query over human mages
	measure("indexed query", "human female mages", func() {
		players.Query(func(txn *column.Txn) error {
			fmt.Printf("-> result = %v\n", txn.With("human", "female", "mage").Count())
			return nil
		})
	}, runs*1000)

	// update everyone
	measure("update", "balance of everyone", func() {
		updates := 0
		players.Query(func(txn *column.Txn) error {
			balance := txn.Float64("balance")
			return txn.Range(func(idx uint32) {
				updates++
				balance.Set(1000.0)
			})
		})
		fmt.Printf("-> updated %v rows\n", updates)
	}, runs)

	// update age of mages
	measure("update", "age of mages", func() {
		updates := 0
		players.Query(func(txn *column.Txn) error {
			age := txn.Int("age")
			return txn.With("mage").Range(func(idx uint32) {
				updates++
				age.Set(99)
			})
		})
		fmt.Printf("-> updated %v rows\n", updates)
	}, runs)
}

// createCollection loads a collection of players
func createCollection(out *column.Collection, amount int) *column.Collection {
	_ = "STUB: not implemented"
	return nil
}

// index for humans

// index for mages

// index for males

// index for females

// Load the data in

// insertPlayers inserts players
func insertPlayers(dst *column.Collection, data []fixtures.Player) error {
	_ = "STUB: not implemented"
	return nil
}

// measure runs a function and measures it
func measure(action, name string, fn func(), iterations int) { _ = "STUB: not implemented"; return }

// Run a few times so the results are more stable

// Silence subsequent runs
