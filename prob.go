package model

import (
	"github.com/aszxqaz/model/kline"
	"github.com/aszxqaz/model/span"
)

func CalculateProb(
	klines []kline.Kline,
	previous []PreviousData,
	target float64,
	secs int,
	volumeSpan span.Span,
	takerSpan span.Span,
	previousPeriod int,
) Prob {
	var (
		cases float64
		hits  float64
		total float64
	)

	IterateWindow(secs, len(klines), previousPeriod, func(a, b int) {
		volume := previous[a].Volume
		taker := previous[a].TakerShare

		if volumeSpan.Includes(volume) &&
			takerSpan.Includes(taker) {
			delta := klines[b].Close - klines[a].Close
			cases++
			switch {
			case target < 0 && delta < target:
				hits++
			case target > 0 && delta > target:
				hits++
			case target == 0 && delta != 0:
				hits += 0.5
			}
		}

		total++
	})

	if cases == 0 {
		return Prob{
			Frequency:   0,
			Probability: 0,
		}
	}

	return Prob{
		Frequency:   cases / total,
		Probability: hits / cases,
	}
}

func IterateWindow(window int, total int, offset int, do func(a int, b int)) {
	for a := offset; a < total-window; a++ {
		do(a, a+window)
	}
}
