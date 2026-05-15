package model

import "math"

func NaturalSumWeight(i, n int) float64 {
	fi := float64(i)
	fn := float64(n)
	return 2 * (fi + 1.0) / (fn * (fn + 1))
}

func QuarticWeight(i, n int) float64 {
	fi := float64(i)
	fn := float64(n)
	return 30 * math.Pow(fi+1, 4) / (fn * (fn + 1) * (2*fn + 1) * (3*fn*fn + 3*fn - 1))
}

func QuadraticWeight(i, n int) float64 {
	fi := float64(i)
	fn := float64(n)
	return 6.0 * (fi + 1.0) * (fi + 1.0) / (fn * (fn + 1) * (2*fn + 1))
}

func CubicWeight(i, n int) float64 {
	fi := float64(i)
	fn := float64(n)
	return 4 * math.Pow(fi+1, 3) / (fn * fn * (fn + 1) * (fn + 1))
}

func GeometricWeight(a float64) func(i, n int) float64 {
	return func(i, n int) float64 {
		fi := float64(i)
		fn := float64(n)
		return a * math.Pow(1-a, fn-1-fi) / (1 - math.Pow(1-a, fn))
	}
}

func EqualWeight(i, n int) float64 {
	return 1.0 / float64(n)
}
