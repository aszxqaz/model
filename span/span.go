package span

import (
	"fmt"
	"strings"

	"github.com/aszxqaz/model/assert"
)

type Span struct {
	Points []float32
}

func New(points ...float32) Span {
	assertInit(points)
	return Span{
		Points: points,
	}
}

func assertInit(points []float32) {
	assert.Panic(len(points) >= 2)

	for i := range len(points) - 1 {
		assert.Panic(points[i] < points[i+1])
	}
}

func (s *Span) String() string {
	r := make([]string, len(s.Points))
	for _, p := range s.Points {
		r = append(r, fmt.Sprintf("%.2f", p))
	}
	return "Span [" + strings.Join(r, ", ") + "]"
}

func (s *Span) Len() int {
	return len(s.Points)
}

func (s *Span) Includes(p float32) bool {
	return p > s.Min() && p < s.Max()
}

func (s *Span) Min() float32 {
	return s.Points[0]
}

func (s *Span) Max() float32 {
	return s.Points[len(s.Points)-1]
}

func (s *Span) Size() int {
	return len(s.Points) - 1
}

func (s *Span) IndexOf(p float32) int {
	for i := range len(s.Points) - 1 {
		if p > s.Points[i] && p < s.Points[i+1] {
			return i
		}
	}
	return -1
}

func (s *Span) Ranges(yield func(s Span) bool) {
	for i := range len(s.Points) - 1 {
		s := New(
			s.Points[i],
			s.Points[i+1],
		)
		if !yield(s) {
			return
		}
	}
}

func (s *Span) SpanFor(p float32) (Span, bool) {
	i := s.IndexOf(p)
	if i == -1 {
		return Span{}, false
	}

	return New(s.Points[i], s.Points[i+1]), true
}
