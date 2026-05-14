package model

import (
	"errors"
	"fmt"
)

func (m *Model) GetProbability(secs int, target int, volume float64, taker float64) (float64, error) {
	if secs < m.Params.SecondsMin || secs > m.Params.SecondsMax {
		msg := fmt.Sprintf("seconds out of range %d..%d", m.Params.SecondsMin, m.Params.SecondsMax)
		return 0, errors.New(msg)
	}

	if target < m.Params.TargetMin || target > m.Params.TargetMax {
		msg := fmt.Sprintf("target out of range %d..%d", m.Params.TargetMin, m.Params.TargetMax)
		return 0, errors.New(msg)
	}

	if target == 0 {
		return 0, errors.New("target is zero")
	}

	volumes := m.Params.Volumes
	volumeMin := volumes[0]
	volumeMax := volumes[len(volumes)-1]

	if volume < volumeMin || volume > volumeMax {
		msg := fmt.Sprintf("volume out of range %.2f .. %.2f", volumeMin, volumeMax)
		return 0, errors.New(msg)
	}

	takers := m.Params.Takers
	takerMin := takers[0]
	takerMax := takers[len(takers)-1]

	if taker < takerMin || taker > takerMax {
		msg := fmt.Sprintf("taker out of range %.2f .. %.2f", takerMin, takerMax)
		return 0, errors.New(msg)
	}

	vi := IndexClosestRange(volume, volumes)
	if vi == -1 {
		return 0, errors.New("volume not found")
	}

	ti := IndexClosestRange(taker, takers)
	if ti == -1 {
		return 0, errors.New("taker not found")
	}

	ri := target - m.Params.TargetMin
	if m.Params.TargetMin < 0 && m.Params.TargetMax > 0 && target > 0 {
		ri--
	}

	si := secs - m.Params.SecondsMin

	prob := m.Probs[si][ri][vi][ti]

	return float64(prob) / 255.0, nil
}

func (m *Model) Size() int {
	return len(m.Probs) * len(m.Probs[0]) * len(m.Probs[0][0]) * len(m.Probs[0][0][0])
}

func (m *Model) Size1() int {
	return len(m.Probs)
}

func (m *Model) Size2() int {
	return len(m.Probs[0])
}

func (m *Model) Size3() int {
	return len(m.Probs[0][0])
}

func (m *Model) Size4() int {
	return len(m.Probs[0][0][0])
}
