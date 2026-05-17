package model

// import (
// 	"errors"
// 	"fmt"
// 	"log/slog"
// 	"math"
// 	"runtime"

// 	"github.com/aszxqaz/model/kline"
// )

// type jobIn struct {
// 	Secs int
// }

// type jobOut struct {
// 	Probs [][][]byte
// 	Secs  int
// }

// func verifyProbes(p Params, b [][][][]byte) error {
// 	if p.SecondsCount() != len(b) {
// 		return errors.New("seconds count mismatch")
// 	}
// 	for i1 := range b {
// 		if len(b[i1]) != p.TargetsCount() {
// 			return errors.New("targets count mismatch")
// 		}
// 		for i2 := range b[i1] {
// 			if len(b[i1][i2]) != p.Volumes.Size() {
// 				return errors.New("volumes count mismatch")
// 			}
// 			for i3 := range b[i1][i2] {
// 				if len(b[i1][i2][i3]) != p.Takers.Size() {
// 					return errors.New("takers count mismatch")
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// func calculateProbs(p Params, klines []kline.Kline) [][][][]byte {
// 	freqs := make([][][][]byte, p.SecondsCount())

// 	slog.Info("Calculating consequtive data...")

// 	conseq := CalcPreviousAll(klines, p.PreviousPeriod, p.WeightFunc)

// 	slog.Info("Assigning jobs...")

// 	out := make(chan jobOut, p.SecondsCount())
// 	in := make(chan jobIn, p.SecondsCount())

// 	for range runtime.NumCPU() {
// 		go worker(p, klines, conseq, in, out)
// 	}

// 	for secs := p.SecondsMin; secs <= p.SecondsMax; secs++ {
// 		in <- jobIn{
// 			Secs: secs,
// 		}
// 	}
// 	close(in)

// 	slog.Info("Waiting for results...")

// 	for i := range p.SecondsCount() {
// 		result := <-out
// 		freqs[result.Secs-p.SecondsMin] = result.Probs
// 		slog.Info(fmt.Sprintf("%d/%d complete", i+1, p.SecondsCount()))
// 	}

// 	return freqs
// }

// func worker(p Params, klines []kline.Kline, conseq []PreviousData, in chan jobIn, out chan jobOut) {
// 	for job := range in {
// 		secs := job.Secs
// 		a1 := [][][]byte{}
// 		for target := p.TargetMin; target <= p.TargetMax; target++ {
// 			if target == 0 {
// 				continue
// 			}
// 			a2 := [][]byte{}
// 			for volumeSpan := range p.Volumes.Ranges {
// 				a3 := []byte{}
// 				for takerSpan := range p.Takers.Ranges {
// 					p := CalculateProb(
// 						klines,
// 						conseq,
// 						float32(target),
// 						secs,
// 						volumeSpan,
// 						takerSpan,
// 						p.PreviousPeriod,
// 					)
// 					a3 = append(a3, fractionToByte(p.Probability))
// 				}
// 				a2 = append(a2, a3)
// 			}
// 			a1 = append(a1, a2)
// 		}
// 		out <- jobOut{
// 			Probs: a1,
// 			Secs:  secs,
// 		}
// 	}
// }

// func fractionToByte(f float32) byte {
// 	return byte(math.Round(f * 255))
// }

// func byteToFraction(b byte) float32 {
// 	return float32(b) / 255
// }
