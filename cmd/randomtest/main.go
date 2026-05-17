package main

import (
	"fmt"
	"math/rand/v2"
)

const (
	advantage float64 = 0
)

func main() {
	var count int

	for range 1_000_000 {
		var balance float64
		for range 272 {
			rnd1 := rand.Float64()
			rnd2 := rand.Float64()
			if rnd2+advantage > rnd1 {
				balance += 1 - rnd1
			} else {
				balance -= rnd1
			}
		}

		if balance >= 16.14 {
			count++
		}
	}

	fmt.Println(1.0 - float64(count)/float64(1_000_000))
}
