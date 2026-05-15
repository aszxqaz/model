package main

import (
	"fmt"

	"github.com/aszxqaz/model/span"
)

func main() {
	s := span.New(1, 2)

	fmt.Println(s.Includes(3))
	fmt.Println(s.Includes(9))
	fmt.Println(s.Includes(1.5))
	fmt.Println(s.Includes(0))
	fmt.Println(s.IndexOf(1.5))
	fmt.Println(s.IndexOf(0))
	fmt.Println(s.IndexOf(5))
	fmt.Println(s.Size())
}
