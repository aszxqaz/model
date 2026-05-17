package main

import (
	"fmt"
	"slices"

	"github.com/aszxqaz/model/kline"
)

// 0 1220
// 0 0.14 0.22 0.3 0.38 0.46 0.55 0.645 0.75 0.86 0.98 1.11 1.25 1.4 1.57 1.77 2 2.22 2.47 2.76 3.1 3.47 3.9 4.4 5 5.76 6.74 8 9.8 12.4 17 28 1220

// 0 0.5 1
// 0 0.22 0.5 0.785 1
// 0 0.102 0.22 0.355 0.5 0.646 0.785 0.905 1
// 0 0.051 0.102 0.158 0.22 0.285 0.355 0.427 0.5 0.573 0.646 0.717 0.785 0.848 0.905 0.954 1

func main() {
	klines, _, err := kline.LoadKlines("C:/Users/Admin/Documents/binance/klines/train/*.csv")
	if err != nil {
		panic(err)
	}

	// period := 31

	// var min float64 = 0.905
	// var max float64 = 1
	// var probe float64 = 0.954

	// count := 0
	// top := 0
	// bottom := 0

	// model.IterateWindow(period, len(klines), 0, func(a, b int) {
	// 	var (
	// 		volume      float64
	// 		takerVolume float64
	// 	)

	// 	for i := a; i < b; i++ {
	// 		w := model.CubicWeight(i-a, period)
	// 		volume += klines[i].Volume * w * float64(period)
	// 		takerVolume += klines[i].TakerVolume * w * float64(period)
	// 	}

	// 	ratio := takerVolume / volume

	// 	if ratio < max && ratio > min {
	// 		count++
	// 		if ratio > probe {
	// 			top++
	// 		} else {
	// 			bottom++
	// 		}
	// 	}
	// })

	// fmt.Println("count", count)
	// fmt.Println("top", top)
	// fmt.Println("bottom", bottom)

	// fmt.Println(float64(top) / float64(bottom))

	slices.SortFunc(klines, func(a, b kline.Kline) int {
		if a.Volume > b.Volume {
			return 1
		}
		return -1
	})

	fmt.Println(klines[len(klines)/2])
}
