package kline

import (
	"encoding/csv"
	"encoding/gob"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"

	"github.com/valyala/fastjson/fastfloat"
)

func LoadKlines(pattern string) ([]Kline, []string, error) {
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

func loadSingle(path string) ([]Kline, error) {
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

	out := make([]Kline, 0, len(records))
	for _, rec := range records {
		close := fastfloat.ParseBestEffort(rec[4])
		volume := fastfloat.ParseBestEffort(rec[5])
		takerVolume := fastfloat.ParseBestEffort(rec[9])

		out = append(out, Kline{
			Close:       close,
			Volume:      volume,
			TakerVolume: takerVolume,
		})
	}
	return out, nil
}

func loadClosePricesFromFiles(paths []string) ([]Kline, error) {
	var all []Kline
	for _, p := range paths {
		part, err := loadSingle(p)
		if err != nil {
			return nil, err
		}
		all = append(all, part...)
	}
	return all, nil
}

func SaveGob(klines []Kline, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)

	if err := encoder.Encode(klines); err != nil {
		return err
	}

	return nil
}

func LoadGob(filename string) ([]Kline, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	var klines []Kline
	if err := decoder.Decode(&klines); err != nil {
		return nil, err
	}

	return klines, nil
}
