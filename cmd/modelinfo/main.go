package main

import (
	"fmt"

	"github.com/aszxqaz/model/model"
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
	fmt.Printf("Targets: %d..%d\n", m.Params.TargetMin, m.Params.TargetMax)
	fmt.Printf("Seconds: %d..%d\n", m.Params.SecondsMin, m.Params.SecondsMax)
	fmt.Println("Volumes: ", m.Params.Volumes)
	fmt.Println("Takers: ", m.Params.Takers)
	fmt.Println(m.Probs)
}
