package model

import (
	"encoding/gob"
	"os"
)

func Save(model *Model, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)

	err = encoder.Encode(model)
	if err != nil {
		return err
	}

	return nil
}
