package main

import (
	"fmt"

	"github.com/aszxqaz/model/kline"
	"github.com/aszxqaz/model/model/calculator"
	"github.com/aszxqaz/model/model/dynparam"
)

func main() {
	load()
}

func save() {
	tklines, _, err := kline.LoadKlines("C:/Users/Admin/Documents/binance/klines/work/*.csv")
	if err != nil {
		panic(err)
	}

	c := calculator.New(calculator.Config{
		Klines: tklines,
		DynamicParams: []calculator.DynamicParam{
			{
				Param: dynparam.NewVolume(dynparam.VolumeDynamicParamConfig{
					Period: 300,
				}),
				Buckets: 32,
			},
			{
				Param: dynparam.NewTaker(dynparam.TakerDynamicParamConfig{
					Period: 300,
				}),
				Buckets: 32,
			},
		},
	})

	err = c.Save("calculator.gob")
	if err != nil {
		panic(err)
	}
}

func load() {

	// c, err := calculator.NewFromFile("calculator.gob")
	// if err != nil {
	// 	panic(err)
	// }

	tklines, _, err := kline.LoadKlines("C:/Users/Admin/Documents/binance/klines/train/*.csv")
	if err != nil {
		panic(err)
	}

	// var (
	// 	min float32 = 99999999
	// 	max float32
	// )

	// for i := range tklines {
	// 	if tklines[i].Close < min {
	// 		min = tklines[i].Close
	// 	}
	// 	if tklines[i].Close > max {
	// 		max = tklines[i].Close
	// 	}
	// }

	// fmt.Println(min)
	// fmt.Println(max)

	c := calculator.New(calculator.Config{
		Klines: tklines,
		DynamicParams: []calculator.DynamicParam{
			{
				Param: dynparam.NewVolume(dynparam.VolumeDynamicParamConfig{
					Period: 400,
				}),
				Buckets: 5,
			},
			{
				Param: dynparam.NewTaker(dynparam.TakerDynamicParamConfig{
					Period: 400,
				}),
				Buckets: 5,
			},
		},
	})

	wklines, _, err := kline.LoadKlines("C:/Users/Admin/Documents/binance/klines/work/*.csv")
	if err != nil {
		panic(err)
	}

	target := 10
	sec := 30

	fmt.Println("Starting")
	_ = wklines
	work := wklines

	cache := make(map[string][2]float32)

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

		delta := work[i+sec].Close - work[i].Close

		if target < 0 && delta < float32(target) || target > 0 && delta > float32(target) {
			ff[0] += (1 - prob.Probability)
		} else {
			ff[0] -= prob.Probability
		}

		ff[1]++
		cache[prob.Key] = ff
	}

	var (
		sum   float32
		count float32
	)

	for k, v := range cache {
		if v[1] == 0 {
			continue
		}
		sum += v[0]
		count += v[1]

		fmt.Printf("%s: %.4f, count = %.0f\n", k, v[0]/v[1], v[1])
	}

	fmt.Printf("Average: %.4f", sum/count)

	fmt.Println("Saving...")
	err = c.Save("calculator.gob")
	if err != nil {
		panic(err)
	}
}
