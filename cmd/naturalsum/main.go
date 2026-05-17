package main

import (
	"fmt"

	"github.com/aszxqaz/model/model"
)

func main() {
	n := 31
	sum := 0.0
	for i := range n {
		w := model.QuadraticWeight(i, n)
		fmt.Printf("%.2f\n", w)
		sum += w
	}

	fmt.Println("sum", sum)
}
