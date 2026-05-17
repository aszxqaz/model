package dynparam

import (
	"github.com/aszxqaz/model/kline"
	"github.com/aszxqaz/model/model"
)

type VolumeDynamicParamConfig struct {
	WeightFunc model.WeightFunc
	Period     int
}

type volumeDynamicParam struct {
	config VolumeDynamicParamConfig
}

func NewVolume(config VolumeDynamicParamConfig) DynamicParam {
	return &volumeDynamicParam{
		config: config,
	}
}

func (v *volumeDynamicParam) Evaluate(klines []kline.Kline) (float64, bool) {
	if len(klines) < v.config.Period {
		return 0, false
	}

	var volume float64

	start := len(klines) - v.config.Period

	for i := start; i < len(klines); i++ {
		volume += klines[i].Volume * v.config.WeightFunc(
			i-start,
			len(klines)-start,
		)
	}

	return volume, true
}
