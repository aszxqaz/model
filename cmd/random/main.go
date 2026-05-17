package main

import (
	"fmt"
	"math/rand"
)

const (
	REAL_PROBABILITY = 0.5
	MY_PROBABILITY   = 0.4
	ITERATIONS_COUNT = 1_000_000
)

func main() {
	err := 0.0

	for range ITERATIONS_COUNT {
		if rand.Float64() < REAL_PROBABILITY {
			// СОБЫТИЕ ПРОИЗОШЛО
			err += (1 - MY_PROBABILITY)
		} else {
			err -= MY_PROBABILITY
		}
	}

	fmt.Println("Expectancy", err/ITERATIONS_COUNT)
}
