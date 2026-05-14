package model

import "time"

type Params struct {
	SecondsMin int       `json:"seconds_min"`
	SecondsMax int       `json:"seconds_max"`
	TargetMin  int       `json:"target_min"`
	TargetMax  int       `json:"target_max"`
	Volumes    []float64 `json:"volumes"`
	Takers     []float64 `json:"takers"`
}

type Model struct {
	Params  Params
	Probs   [][][][]byte
	Files   []string
	Created time.Time
}

type kline struct {
	Close       float64
	Volume      float64
	TakerVolume float64
}

func (p *Params) SecondsCount() int {
	return p.SecondsMax - p.SecondsMin + 1
}

func (p *Params) TargetsCount() int {
	if p.TargetMin < 0 && p.TargetMax > 0 {
		return p.TargetMax - p.TargetMin
	}

	return p.TargetMax - p.TargetMin + 1
}

func (p *Params) VolumesCount() int {
	return len(p.Volumes) - 1
}

func (p *Params) TakersCount() int {
	return len(p.Takers) - 1
}
