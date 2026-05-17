package calculator

func makeRanges(length int, count int) []int {
	if length <= count {
		panic("length <= count")
	}
	first := 0
	second := length - 1
	rest := []int{first}
	i := 0
	j := count
	for j-1 > 0 {
		i += second / count
		rest = append(rest, i)
		j--
	}

	return append(rest, []int{second}...)
}
