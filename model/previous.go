package model

// type PreviousData struct {
// 	Volume     float32
// 	TakerShare float32
// }

// func CalcPreviousAll(klines []kline.Kline, period int, wght func(i, n int) float32) []PreviousData {
// 	data := make([]PreviousData, len(klines))

// 	for k := period; k < len(klines); k++ {
// 		data[k] = CalcPreviousSingle(klines[k-period:k], wght)
// 	}

// 	return data
// }

// func CalcPreviousSingle(klines []kline.Kline, wght func(i, n int) float32) PreviousData {
// 	var (
// 		volume      float32
// 		takerVolume float32
// 	)

// 	flen := float32(len(klines))

// 	for i := range klines {
// 		w := wght(i, len(klines))
// 		volume += klines[i].Volume * w * flen
// 		takerVolume += klines[i].TakerVolume * w * flen
// 	}

// 	return PreviousData{
// 		Volume:     volume,
// 		TakerShare: takerVolume / volume,
// 	}
// }
