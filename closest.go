package model

func IndexClosestRange(target float64, candidates []float64) int {
	for i := 0; i < len(candidates)-1; i++ {
		if target >= candidates[i] && target <= candidates[i+1] {
			return i
		}
	}
	return -1
}
