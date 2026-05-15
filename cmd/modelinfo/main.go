package main

import (
	"fmt"

	"github.com/aszxqaz/model"
)

func main() {
	m, err := model.Load("testdata/models/model.gob")
	if err != nil {
		panic(err)
	}

	fmt.Println("Files:")
	fmt.Println("\t" + m.Files[0])
	fmt.Println("\t ... ")
	fmt.Println("\t" + m.Files[len(m.Files)-1])
	fmt.Printf("Model size: %dK\n", m.Size()/1000)
	fmt.Printf("Targets: %d..%d\n", m.Params.TargetMin, m.Params.TargetMax)
	fmt.Printf("Seconds: %d..%d\n", m.Params.SecondsMin, m.Params.SecondsMax)
	fmt.Println("Volumes: ", m.Params.Volumes)
	fmt.Println("Takers: ", m.Params.Takers)
	fmt.Println("Size1: ", m.Size1())
	fmt.Println("Size2: ", m.Size2())
	fmt.Println("Size3: ", m.Size3())
	fmt.Println("Size4: ", m.Size4())
	fmt.Println(m.Probs)
}
