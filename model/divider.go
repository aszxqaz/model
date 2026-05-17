package model

// func GetVolumesSpaced(previous []PreviousData, count int) []float64 {
// 	slices.SortFunc(previous, func(a, b PreviousData) int {
// 		if a.Volume > b.Volume {
// 			return 1
// 		}
// 		return -1
// 	})

// 	indeces := MakeRanges(len(previous), count)

// 	volumes := []float64{}

// 	for _, i := range indeces {
// 		volumes = append(volumes, previous[i].Volume)
// 	}

// 	return volumes
// }

// func GetTakersSpaced(previous []PreviousData, count int) []float64 {
// 	slices.SortFunc(previous, func(a, b PreviousData) int {
// 		if a.TakerShare > b.TakerShare {
// 			return 1
// 		}
// 		return -1
// 	})

// 	indeces := MakeRanges(len(previous), count)

// 	takers := []float64{}

// 	for _, i := range indeces {
// 		takers = append(takers, previous[i].TakerShare)
// 	}

// 	return takers
// }
