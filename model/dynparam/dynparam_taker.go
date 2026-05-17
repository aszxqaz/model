package dynparam

import (
	"encoding/gob"

	"github.com/aszxqaz/model/kline"
	"github.com/aszxqaz/model/model"
)

func init() {
	gob.Register(TakerDynamicParam{})
}

type TakerDynamicParamConfig struct {
	Period int
}

type TakerDynamicParam struct {
	Config TakerDynamicParamConfig
}

func NewTaker(config TakerDynamicParamConfig) DynamicParam {
	return TakerDynamicParam{
		Config: config,
	}
}

func (v TakerDynamicParam) Evaluate(klines []kline.Kline) (float32, bool) {
	if len(klines) < v.Config.Period {
		return 0, false
	}

	var (
		takerVolume float32
		volume      float32
	)

	start := len(klines) - v.Config.Period

	for i := start; i < len(klines); i++ {
		w := float32(model.QuadraticWeight(i-start, len(klines)-start))
		volume += klines[i].Volume * w
		takerVolume += klines[i].TakerVolume * w

	}

	return volume / takerVolume, true
}
