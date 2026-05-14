package main

import (
	"fmt"

	"github.com/aszxqaz/model"
)

func main() {
	s := []float64{0, 1, 2, 3, 4, 5}
	fmt.Println(model.IndexClosestRange(5.5, s))
	fmt.Println(len(s))
}
