package calculator

import (
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"slices"

	"github.com/aszxqaz/model/kline"
	"github.com/aszxqaz/model/model/dynparam"
	"github.com/aszxqaz/model/span"
)

type DynamicParam struct {
	Param   dynparam.DynamicParam
	Buckets int
}

type Config struct {
	Klines        []kline.Kline
	DynamicParams []DynamicParam
}

type Calculator struct {
	Config  Config
	Dynvals [][]float32
	Dynnils [][]bool
	Spans   []span.Span
	Probs   map[string]Prob
}

func New(c Config) *Calculator {
	calc := &Calculator{
		Config: c,
		Probs:  make(map[string]Prob),
	}
	calc.prepareDyns()
	calc.prepareSpans()
	return calc
}

func (c *Calculator) prepareDyns() {
	dynvals := make([][]float32, len(c.Config.DynamicParams))
	dynnils := make([][]bool, len(c.Config.DynamicParams))
	fmt.Println("Preparing dyns...")
	for i := range dynvals {
		if c.Config.DynamicParams[i].Buckets > 0 {
			vals := make([]float32, len(c.Config.Klines))
			nils := make([]bool, len(c.Config.Klines))
			for k := 0; k < len(c.Config.Klines); k++ {
				val, ok := c.Config.DynamicParams[i].Param.Evaluate(c.Config.Klines[:k])
				if !ok {
					nils[k] = true
				} else {
					vals[k] = val
				}
			}
			dynvals[i] = vals
			dynnils[i] = nils
		}
		fmt.Printf("Prepared %d of %d\n", i+1, len(dynnils))
	}
	c.Dynvals = dynvals
	c.Dynnils = dynnils
}

func (c *Calculator) prepareSpans() {
	spans := make([]span.Span, len(c.Config.DynamicParams))
	fmt.Println("Preparing spans...")
	for i := range spans {
		if c.Config.DynamicParams[i].Buckets > 0 {
			vals := slices.Clone(c.Dynvals[i])
			deletions := 0
			for k := 0; k < len(c.Dynvals[i]); k++ {
				if c.Dynnils[i][k] {
					vals = append(vals[:k-deletions], vals[k-deletions+1:]...)
					deletions++
				}
			}

			slices.Sort(vals)
			indeces := makeRanges(len(vals), c.Config.DynamicParams[i].Buckets)
			points := make([]float32, len(indeces))
			for j := range points {
				points[j] = vals[indeces[j]]
			}
			spans[i] = span.New(points...)
		}
		fmt.Printf("Prepared %d of %d\n", i+1, len(spans))
	}

	c.Spans = spans
}

type Prob struct {
	Probability float32
	Frequency   float32
	Key         string
}

func (c *Calculator) GetBuckets(target int, seconds int, prev []kline.Kline) (string, []int, error) {
	if target == 0 {
		return "", nil, errors.New("target is zero")
	}

	key := fmt.Sprintf("%d-%d", target, seconds)
	indeces := make([]int, len(c.Config.DynamicParams))
	for i := range c.Config.DynamicParams {
		val, ok := c.Config.DynamicParams[i].Param.Evaluate(prev)
		if ok {
			d := c.Spans[i].IndexOf(val)
			if d == -1 {
				return "", nil, errors.New("val not found in span")
			}
			key = fmt.Sprintf("%s-%d", key, d)
			indeces[i] = d
		} else {
			return "", nil, errors.New("can not eval dynval for klines")
		}
	}

	return key, indeces, nil
}

func (c *Calculator) CountProb(target int, seconds int, indeces []int, key string) (Prob, error) {
	prob, ok := c.Probs[key]
	if ok {
		return prob, nil
	}

	var (
		cases float32
		hits  float32
		total float32
	)

outer:
	for k := 0; k < len(c.Config.Klines)-seconds; k++ {
		total++

		for i := range c.Config.DynamicParams {
			if c.Dynnils[i][k] {
				continue outer
			}

			d := c.Spans[i].IndexOf(c.Dynvals[i][k])
			if d == -1 {
				continue outer
			}

			if indeces[i] != d {
				continue outer
			}
		}

		cases++
		delta := c.Config.Klines[k+seconds].Close - c.Config.Klines[k].Close
		switch {
		case target < 0 && delta < float32(target):
			hits++
		case target > 0 && delta > float32(target):
			hits++
		}
	}

	if cases == 0 {
		prob = Prob{
			Frequency:   0,
			Probability: 0,
			Key:         key,
		}
	} else {
		prob = Prob{
			Frequency:   cases / total,
			Probability: hits / cases,
			Key:         key,
		}
	}

	c.Probs[key] = prob
	return prob, nil
}

func NewFromFile(filename string) (*Calculator, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	var c Calculator
	if err := decoder.Decode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Calculator) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)

	if err := encoder.Encode(c); err != nil {
		return err
	}

	return nil
}
