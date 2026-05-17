package kline

import (
	"encoding/gob"
	"os"
)

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
