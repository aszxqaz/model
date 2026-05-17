package dynparam

import "github.com/aszxqaz/model/kline"

type DynamicParam interface {
	Evaluate(klines []kline.Kline) (float32, bool)
}
