package model

type Result struct {
	Total float64
	Prob  float64
}

func calculateProb(
	klines []kline,
	data []consequtiveData,
	target float64,
	secs int,
	minVolume float64,
	maxVolume float64,
	minTaker float64,
	maxTaker float64,
) Result {
	var (
		total float64
		count float64
	)

	iterateWindow(secs, len(klines), 60, func(a, b int) {
		volume := data[a].Volume
		taker := data[a].Taker

		if volume > minVolume && volume < maxVolume && taker > minTaker && taker < maxTaker {
			delta := klines[b].Close - klines[a].Close
			total++
			switch {
			case target < 0 && delta < target:
				count++
			case target > 0 && delta > target:
				count++
			case target == 0 && delta != 0:
				count += 0.5
			}
		}

	})

	if total == 0 {
		return Result{
			Total: 0,
			Prob:  0,
		}
	}
	return Result{
		Total: total,
		Prob:  count / total,
	}
}

func iterateWindow(window int, total int, offset int, do func(a int, b int)) {
	for a := offset; a < total-window; a++ {
		do(a, a+window)
	}
}
