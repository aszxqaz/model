package model

import (
	"slices"
)

func GetVolumesSpaced(previous []PreviousData, size int) []float64 {
	slices.SortFunc(previous, func(a, b PreviousData) int {
		if a.Volume > b.Volume {
			return 1
		}
		return -1
	})

	step := len(previous) / (size + 1)

	volumes := []float64{}

	for i := range size + 1 {
		volumes = append(volumes, previous[i*step].Volume)
	}

	return volumes
}

func GetTakersSpaced(previous []PreviousData, size int) []float64 {
	slices.SortFunc(previous, func(a, b PreviousData) int {
		if a.TakerShare > b.TakerShare {
			return 1
		}
		return -1
	})

	step := len(previous) / (size + 1)

	takers := []float64{}

	for i := range size + 1 {
		takers = append(takers, previous[i*step].TakerShare)
	}

	return takers
}

// 1 2 3 4 5 6 7 8 9  // 4
// 1   3   5   7   9
// 0   2   4   6   8
