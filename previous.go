package model

import "github.com/aszxqaz/model/kline"

type PreviousData struct {
	Volume     float64
	TakerShare float64
}

func CalcPreviousAll(klines []kline.Kline, period int, wght func(i, n int) float64) []PreviousData {
	data := make([]PreviousData, len(klines))

	for k := period; k < len(klines); k++ {
		data[k] = CalcPreviousSingle(klines[k-period:k], wght)
	}

	return data
}

func CalcPreviousSingle(klines []kline.Kline, wght func(i, n int) float64) PreviousData {
	var (
		volume      float64
		takerVolume float64
	)

	flen := float64(len(klines))

	for i := range klines {
		w := wght(i, len(klines))
		volume += klines[i].Volume * w * flen
		takerVolume += klines[i].TakerVolume * w * flen
	}

	return PreviousData{
		Volume:     volume,
		TakerShare: takerVolume / volume,
	}
}
