package dynparam

import (
	"encoding/gob"
	"math"

	"github.com/aszxqaz/model/kline"
	"github.com/aszxqaz/model/model"
)

func init() {
	gob.Register(VolumeDynamicParam{})
}

type VolumeDynamicParamConfig struct {
	Period int
}

type VolumeDynamicParam struct {
	Config VolumeDynamicParamConfig
}

func NewVolume(config VolumeDynamicParamConfig) DynamicParam {
	return VolumeDynamicParam{
		Config: config,
	}
}

func (v VolumeDynamicParam) Evaluate(klines []kline.Kline) (float32, bool) {
	if len(klines) < v.Config.Period {
		return 0, false
	}

	var volume float32

	start := len(klines) - v.Config.Period

	for i := start; i < len(klines); i++ {
		volume += klines[i].Volume * float32(model.CubicWeight(
			i-start,
			len(klines)-start,
		))
	}

	return volume, true

	// sum := 0.0
	// start := len(klines) - v.config.Period

	// power := -1.0

	// for i := start; i < len(klines); i++ {
	// 	value := math.Max(0.001, klines[i].Volume)
	// 	sum += math.Pow(value, power)
	// }

	// flen := float32(len(klines))
	// return math.Pow(
	// 	sum/flen, 1/power,
	// ), true

}

func pmean(klines []kline.Kline, power float64) float64 {
	sum := 0.0
	for _, value := range klines {
		value := math.Max(0.001, float64(value.Volume))
		sum += math.Pow(value, power)
	}

	flen := float64(len(klines))
	return math.Pow(
		sum/flen, 1/power,
	)
}
