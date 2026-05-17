package main

import (
	"fmt"
	"math"
)

func main() {
	values := []float64{0.001, 0.002, 0.5, 0.002, 0.001}

	fmt.Println("Values:", values)

	p := -1.0
	fmt.Printf("Ariphmetic mean: %f\n", pmean(values, 1))
	fmt.Printf("Power mean (p = %.4f): %.4f\n", p, pmean(values, p))

	values = []float64{10, 10, 10, 10, 10, 10}
	fmt.Println("\nValues:", values)

	p = -1.0
	fmt.Printf("Ariphmetic mean: %.2f\n", pmean(values, 1))
	fmt.Printf("Power mean (p = %.4f): %.2f\n", p, pmean(values, p))
}

func pmean(values []float64, power float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += math.Pow(value, power)
	}

	flen := float64(len(values))
	return math.Pow(
		sum/flen, 1/power,
	)
}
