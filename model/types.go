package model

import (
	"time"

	"github.com/aszxqaz/model/span"
)

type Params struct {
	SecondsMin     int       `json:"seconds_min"`
	SecondsMax     int       `json:"seconds_max"`
	TargetMin      int       `json:"target_min"`
	TargetMax      int       `json:"target_max"`
	PreviousPeriod int       `json:"previous_period"`
	Volumes        span.Span `json:"volumes"`
	Takers         span.Span `json:"takers"`
	WeightFunc     func(i, n int) float32
}

type Model struct {
	Params  Params
	Probs   [][][][]byte
	Files   []string
	Created time.Time
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
