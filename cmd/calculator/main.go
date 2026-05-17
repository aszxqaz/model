package main

import (
	"fmt"

	"github.com/aszxqaz/model/kline"
	"github.com/aszxqaz/model/model"
	"github.com/aszxqaz/model/model/calculator"
	"github.com/aszxqaz/model/model/dynparam"
)

func main() {
	tklines, _, err := kline.LoadKlines("C:/Users/Admin/Documents/binance/klines/train/*.csv")
	if err != nil {
		panic(err)
	}

	c := calculator.New(calculator.Config{
		Klines: tklines,
		DynamicParams: []calculator.DynamicParam{
			{
				Param: dynparam.NewVolume(dynparam.VolumeDynamicParamConfig{
					WeightFunc: model.QuadraticWeight,
					Period:     30,
				}),
				Buckets: 5,
			},
			{
				Param: dynparam.NewTaker(dynparam.TakerDynamicParamConfig{
					WeightFunc: model.QuadraticWeight,
					Period:     30,
				}),
				Buckets: 5,
			},
		},
	})

	wklines, _, err := kline.LoadKlines("C:/Users/Admin/Documents/binance/klines/work/*.csv")
	if err != nil {
		panic(err)
	}

	// now := time.Now()

	// prob, err := c.CountProb(100, 30, wklines[:29])
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Probability", prob.Probability)
	// fmt.Println("Frequency", prob.Frequency)
	// fmt.Printf("Done in %d ms.\n", time.Since(now).Milliseconds())

	// score := 0.0
	target := 50
	sec := 90
	progress := 0

	fmt.Println("Starting")
	_ = wklines
	work := wklines

	cache := make(map[string][2]float64)

	for i := 0; i < len(work)-sec-1; i++ {
		key, indeces, err := c.GetBuckets(target, sec, work[:i])
		if err != nil {
			continue
		}

		prob, err := c.CountProb(target, sec, indeces, key)
		if err != nil {
			continue
		}

		ff, ok := cache[prob.Key]
		if !ok {
			cache[prob.Key] = ff
		}

		if work[i+sec].Close-work[i].Close >= float64(target) {
			ff[0] += (1 - prob.Probability)
		} else {
			ff[0] -= prob.Probability
		}

		ff[1]++
		cache[prob.Key] = ff

		if int(100*float64(i)/float64(len(work)-sec-1)) != progress {
			progress = int(100 * float64(i) / float64(len(work)-sec-1))
			fmt.Println(progress)
		}
	}

	for k, v := range cache {
		if v[1] == 0 {
			continue
		}

		fmt.Printf("%s: %.2f, count = %.0f\n", k, v[0]/v[1], v[1])
	}
}
