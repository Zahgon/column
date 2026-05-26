// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"fmt"

	"github.com/kelindar/xxrand"
)

func main() {
	amount := 50000
	cache := New()

	measure("insert", fmt.Sprintf("%v rows", amount), func() {
		for i := 0; i < amount; i++ {
			key := fmt.Sprintf("user_%d", i)
			val := fmt.Sprintf("Hi, User %d", i)
			cache.Set(key, val)

			if (i+1)%10000 == 0 {
				fmt.Printf("-> inserted %v rows\n", i+1)
			}
		}
	}, 1)

	key := fmt.Sprintf("user_%d", xxrand.Intn(amount))
	measure("query", key, func() {
		xxrand.Intn(amount)
		fmt.Println(cache.Get(key))
	}, 100000)
}

func measure(action, name string, fn func(), iterations int) { _ = "STUB: not implemented"; return }

// Run a few times so the results are more stable

// Silence subsequent runs
