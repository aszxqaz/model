package dynparam

import "github.com/aszxqaz/model/kline"

type DynamicParam interface {
	Evaluate(klines []kline.Kline) (float64, bool)
}
