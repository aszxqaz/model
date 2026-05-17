package model

// type Prob struct {
// 	Probability float64
// 	Frequency   float64
// }

// func (m *Model) GetProbability(secs int, target int, volume float64, taker float64) (*Prob, error) {
// 	if secs < m.Params.SecondsMin || secs > m.Params.SecondsMax {
// 		msg := fmt.Sprintf("seconds out of range %d..%d", m.Params.SecondsMin, m.Params.SecondsMax)
// 		return nil, errors.New(msg)
// 	}

// 	if target < m.Params.TargetMin || target > m.Params.TargetMax {
// 		msg := fmt.Sprintf("target out of range %d..%d", m.Params.TargetMin, m.Params.TargetMax)
// 		return nil, errors.New(msg)
// 	}

// 	if target == 0 {
// 		return nil, errors.New("target is zero")
// 	}

// 	volumes := m.Params.Volumes
// 	if !volumes.Includes(volume) {
// 		msg := fmt.Sprintf("volume out of range %.2f .. %.2f", volumes.Min(), volumes.Max())
// 		return nil, errors.New(msg)
// 	}

// 	takers := m.Params.Takers
// 	if !takers.Includes(taker) {
// 		msg := fmt.Sprintf("taker out of range %.2f .. %.2f", takers.Min(), takers.Max())
// 		return nil, errors.New(msg)
// 	}

// 	vi := volumes.IndexOf(volume)
// 	if vi == -1 {
// 		return nil, errors.New("volume not found")
// 	}

// 	ti := takers.IndexOf(taker)
// 	if ti == -1 {
// 		return nil, errors.New("taker not found")
// 	}

// 	ri := target - m.Params.TargetMin
// 	if m.Params.TargetMin < 0 && m.Params.TargetMax > 0 && target > 0 {
// 		ri--
// 	}

// 	si := secs - m.Params.SecondsMin

// 	p := m.Probs[si][ri][vi][ti]

// 	// fmt.Println(p[1])

// 	return &Prob{
// 		Probability: byteToFraction(p),
// 	}, nil
// }

// func (m *Model) Size() int {
// 	return len(m.Probs) * len(m.Probs[0]) * len(m.Probs[0][0]) * len(m.Probs[0][0][0])
// }

// func (m *Model) Size1() int {
// 	return len(m.Probs)
// }

// func (m *Model) Size2() int {
// 	return len(m.Probs[0])
// }

// func (m *Model) Size3() int {
// 	return len(m.Probs[0][0])
// }

// func (m *Model) Size4() int {
// 	return len(m.Probs[0][0][0])
// }
