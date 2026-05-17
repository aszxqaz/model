package model

import (
	"encoding/gob"
	"os"
)

func Load(path string) (*Model, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	var model Model
	err = decoder.Decode(&model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
