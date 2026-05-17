package dynparam

import (
	"github.com/aszxqaz/model/kline"
	"github.com/aszxqaz/model/model"
)

type TakerDynamicParamConfig struct {
	WeightFunc model.WeightFunc
	Period     int
}

type TakerDynamicParam struct {
	config TakerDynamicParamConfig
}

func NewTaker(config TakerDynamicParamConfig) DynamicParam {
	return &TakerDynamicParam{
		config: config,
	}
}

func (v *TakerDynamicParam) Evaluate(klines []kline.Kline) (float64, bool) {
	if len(klines) < v.config.Period {
		return 0, false
	}

	var (
		takerVolume float64
		volume      float64
	)

	start := len(klines) - v.config.Period

	for i := start; i < len(klines); i++ {
		w := v.config.WeightFunc(i-start, len(klines)-start)
		volume += klines[i].Volume * w
		takerVolume += klines[i].TakerVolume * w

	}

	return volume / takerVolume, true
}
