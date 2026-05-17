package main

import (
	"github.com/aszxqaz/model/model"
)

type Config struct {
	Params model.Params `json:"params"`
	Klines string       `json:"klines"`
}

func main() {
	// path := flag.String("config", "model.json", "config file")
	// flag.Parse()

	// file, err := os.Open(*path)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// decoder := json.NewDecoder(file)

	// var config Config
	// err = decoder.Decode(&config)
	// if err != nil {
	// 	panic(err)
	// }

	// if config.Params.SecondsMin == 0 {
	// 	panic("seconds_min not set")
	// }

	// if config.Params.SecondsMax == 0 {
	// 	panic("seconds_max not set")
	// }

	// if config.Params.TargetMin == 0 {
	// 	panic("target_min not set")
	// }

	// if config.Params.TargetMax == 0 {
	// 	panic("target_max not set")
	// }

	// if len(config.Params.Takers) < 2 {
	// 	panic("takers length < 2")
	// }

	// if len(config.Params.Volumes) < 2 {
	// 	panic("volumes length < 2")
	// }

	// m, err := model.Generate(config.Klines, config.Params)
	// if err != nil {
	// 	panic(err)
	// }

	// err = model.Save(m, "model.gob")
	// if err != nil {
	// 	panic(err)
	// }
}
