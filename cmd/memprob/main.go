package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aszxqaz/model"
	"github.com/aszxqaz/model/kline"
	"github.com/aszxqaz/model/span"
)

func main() {
	klines, _, err := kline.LoadKlines("C:/Users/Admin/Documents/binance/klines/*.csv")
	if err != nil {
		panic(err)
	}

	p := model.Params{
		PreviousPeriod: 30,
		WeightFunc:     model.CubicWeight,
	}

	previous := model.CalcPreviousAll(klines, p.PreviousPeriod, p.WeightFunc)

	p.Volumes = span.New(model.GetVolumesSpaced(previous, 4)...)
	p.Takers = span.New(model.GetTakersSpaced(previous, 4)...)

	Run(p, func(i Input) error {
		now := time.Now()

		volume, ok := p.Volumes.SpanFor(i.Volume)
		if !ok {
			return errors.New("volume not found")
		}

		taker, ok := p.Takers.SpanFor(i.Taker)
		if !ok {
			return errors.New("taker not found")
		}

		prob := model.CalculateProb(
			klines,
			previous,
			i.Target,
			i.Seconds,
			volume,
			taker,
			p.PreviousPeriod,
		)

		fmt.Printf("\tProbability: %.2f%%\n", prob.Probability*100)
		fmt.Printf("\tDone in %d ms.\n", time.Since(now).Milliseconds())

		return nil
	})
}

type Input struct {
	Target  float64
	Seconds int
	Volume  float64
	Taker   float64
}

func Run(params model.Params, doFunc func(input Input) error) {
	fmt.Println("\nVolumes: ", params.Volumes)
	fmt.Println("\tSize: ", params.Volumes.Size())
	fmt.Println("Takers: ", params.Takers)
	fmt.Println("\tSize: ", params.Takers.Size())
	fmt.Println("---")
	fmt.Fprintln(os.Stderr, "enter: <target> <seconds> <volume> <taker>")

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		fs := strings.Fields(line)
		if len(fs) != 4 {
			fmt.Fprintln(os.Stderr, "bad input; expected: <target> <seconds> <volume> <taker>")
			continue
		}

		target, err := strconv.ParseFloat(fs[0], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "bad target:", err)
			continue
		}

		seconds, err := strconv.Atoi(fs[1])
		if err != nil || seconds < 1 {
			fmt.Fprintln(os.Stderr, "bad seconds (must be int >= 1)")
			continue
		}

		volume, err := strconv.ParseFloat(fs[2], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "bad volume:", err)
			continue
		}

		taker, err := strconv.ParseFloat(fs[3], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "bad taker:", err)
			continue
		}

		err = doFunc(Input{
			Target:  target,
			Seconds: seconds,
			Volume:  volume,
			Taker:   taker,
		})

		if err != nil {
			fmt.Println("\tError:" + err.Error())
		}
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read stdin:", err)
		os.Exit(1)
	}
}
