package dynparam

import (
	"encoding/gob"

	"github.com/aszxqaz/model/kline"
	"github.com/aszxqaz/model/model"
)

func init() {
	gob.Register(TakerVolumeDynamicParam{})
}

type TakerVolumeDynamicParamConfig struct {
	Period int
}

type TakerVolumeDynamicParam struct {
	Config TakerVolumeDynamicParamConfig
}

func NewTakerVolume(config TakerVolumeDynamicParamConfig) DynamicParam {
	return TakerVolumeDynamicParam{
		Config: config,
	}
}

func (v TakerVolumeDynamicParam) Evaluate(klines []kline.Kline) (float32, bool) {
	if len(klines) < v.Config.Period {
		return 0, false
	}

	var (
		takerVolume float32
	)

	start := len(klines) - v.Config.Period

	for i := start; i < len(klines); i++ {
		var w float32 = float32(model.CubicWeight(i-start, len(klines)-start))
		takerVolume += klines[i].TakerVolume * w

	}

	return takerVolume, true
}
