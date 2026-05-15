package model

import (
	"slices"
)

func GetVolumesSpaced(previous []PreviousData, depth int) []float64 {
	slices.SortFunc(previous, func(a, b PreviousData) int {
		if a.Volume > b.Volume {
			return 1
		}
		return -1
	})

	indeces := Divide(0, len(previous)-1, depth)

	volumes := []float64{}

	for _, i := range indeces {
		volumes = append(volumes, previous[i].Volume)
	}

	return volumes
}

func GetTakersSpaced(previous []PreviousData, depth int) []float64 {
	slices.SortFunc(previous, func(a, b PreviousData) int {
		if a.TakerShare > b.TakerShare {
			return 1
		}
		return -1
	})

	indeces := Divide(0, len(previous)-1, depth)

	takers := []float64{}

	for _, i := range indeces {
		takers = append(takers, previous[i].TakerShare)
	}

	return takers
}

func Divide(left, right, depth int) []int {
	if depth == 0 {
		return []int{left, right}
	}

	mid := (left + right) / 2

	leftPart := Divide(left, mid, depth-1)
	rightPart := Divide(mid, right, depth-1)

	// remove duplicate midpoint
	return append(leftPart[:len(leftPart)-1], rightPart...)
}
