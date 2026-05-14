package main

import (
	"github.com/aszxqaz/model"
)

func main() {
	m, err := model.Generate("testdata/klines/*.csv", model.Params{
		SecondsMin: 50,
		SecondsMax: 50,
		TargetMin:  10,
		TargetMax:  10,
		Volumes:    []float64{0, 1, 2, 5, 10, 50},
		Takers:     []float64{0, 0.25, 0.5, 0.75, 1},
	})
	if err != nil {
		panic(err)
	}

	err = model.Save(m, "testdata/models/model.gob")
	if err != nil {
		panic(err)
	}
}
