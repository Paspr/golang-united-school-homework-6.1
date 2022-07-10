package golang_united_school_homework

import (
	"math"
)

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	// assume the trianlge is equilateral
	return 3 * t.Side
}

func (t Triangle) CalcArea() float64 {
	// assume the trianlge is equilateral
	return (math.Sqrt(3) * t.Side * t.Side) / 4
}
