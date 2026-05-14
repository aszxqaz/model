package model

type consequtiveData struct {
	Volume float64
	Taker  float64
}

func getConsequtiveData(klines []kline) []consequtiveData {
	data := make([]consequtiveData, len(klines))

	for k := 60; k < len(klines); k++ {
		var (
			volume float64
			taker  float64
		)

		for i := k - 60; i < k; i++ {
			volume += klines[i].Volume
			taker += klines[i].TakerVolume
		}

		data[k] = consequtiveData{
			Volume: volume,
			Taker:  taker / volume,
		}
	}

	return data
}
