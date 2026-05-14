package model

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"

	"github.com/valyala/fastjson/fastfloat"
)

func loadKlines(pattern string) ([]kline, []string, error) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, nil, err
	}
	if len(matches) == 0 {
		return nil, nil, errors.New("no CSV files found")
	}

	sort.Strings(matches)
	prices, err := loadClosePricesFromFiles(matches)
	if err != nil {
		return nil, nil, err
	}

	return prices, matches, nil
}

func loadSingle(path string) ([]kline, error) {
	slog.Info(fmt.Sprintf("Loading klines from %s", path))
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}

	out := make([]kline, 0, len(records))
	for _, rec := range records {
		close := fastfloat.ParseBestEffort(rec[4])
		volume := fastfloat.ParseBestEffort(rec[5])
		takerVolume := fastfloat.ParseBestEffort(rec[9])

		out = append(out, kline{
			Close:       close,
			Volume:      volume,
			TakerVolume: takerVolume,
		})
	}
	return out, nil
}

func loadClosePricesFromFiles(paths []string) ([]kline, error) {
	var all []kline
	for _, p := range paths {
		part, err := loadSingle(p)
		if err != nil {
			return nil, err
		}
		all = append(all, part...)
	}
	return all, nil
}
