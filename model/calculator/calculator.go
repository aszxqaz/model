package calculator

import (
	"errors"
	"fmt"
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
	c       Config
	Dynvals [][]float64
	Dynnils [][]bool
	Spans   []span.Span
	probs   map[string]Prob
}

func New(c Config) *Calculator {
	calc := &Calculator{
		c:     c,
		probs: make(map[string]Prob),
	}
	calc.prepareDyns()
	calc.prepareSpans()
	return calc
}

func (c *Calculator) prepareDyns() {
	dynvals := make([][]float64, len(c.c.DynamicParams))
	dynnils := make([][]bool, len(c.c.DynamicParams))
	for i := range dynvals {
		if c.c.DynamicParams[i].Buckets > 0 {
			vals := make([]float64, len(c.c.Klines))
			nils := make([]bool, len(c.c.Klines))
			for k := 0; k < len(c.c.Klines); k++ {
				val, ok := c.c.DynamicParams[i].Param.Evaluate(c.c.Klines[:k])
				if !ok {
					nils[k] = true
				} else {
					vals[k] = val
				}
			}
			dynvals[i] = vals
			dynnils[i] = nils
		}
	}
	c.Dynvals = dynvals
	c.Dynnils = dynnils
}

func (c *Calculator) prepareSpans() {
	spans := make([]span.Span, len(c.c.DynamicParams))

	for i := range spans {
		if c.c.DynamicParams[i].Buckets > 0 {
			vals := slices.Clone(c.Dynvals[i])
			deletions := 0
			for k := 0; k < len(c.Dynvals[i]); k++ {
				if c.Dynnils[i][k] {
					vals = append(vals[:k-deletions], vals[k-deletions+1:]...)
					deletions++
				}
			}

			slices.Sort(vals)
			indeces := makeRanges(len(vals), c.c.DynamicParams[i].Buckets)
			points := make([]float64, len(indeces))
			for j := range points {
				points[j] = vals[indeces[j]]
			}
			spans[i] = span.New(points...)
		}
	}

	c.Spans = spans
}

type Prob struct {
	Probability float64
	Frequency   float64
	Key         string
}

func (c *Calculator) GetBuckets(target int, seconds int, prev []kline.Kline) (string, []int, error) {
	if target == 0 {
		return "", nil, errors.New("target is zero")
	}

	key := fmt.Sprintf("%d-%d", target, seconds)
	indeces := make([]int, len(c.c.DynamicParams))
	for i := range c.c.DynamicParams {
		val, ok := c.c.DynamicParams[i].Param.Evaluate(prev)
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
	prob, ok := c.probs[key]
	if ok {
		return prob, nil
	}

	var (
		cases float64
		hits  float64
		total float64
	)

outer:
	for k := 0; k < len(c.c.Klines)-seconds; k++ {
		total++

		for i := range c.c.DynamicParams {
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
		delta := c.c.Klines[k+seconds].Close - c.c.Klines[k].Close
		switch {
		case target < 0 && delta < float64(target):
			hits++
		case target > 0 && delta > float64(target):
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

	c.probs[key] = prob
	return prob, nil
}
