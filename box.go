package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	return errors.New("capacity is exceeded")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("index out of range")
	}

	for k, v := range b.shapes {
		if i == k {
			return v, nil
		}
	}

	return nil, errors.New("such shape doesn't exist")

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, error := b.GetByIndex(i)
	if error != nil {
		return nil, error
	}

	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return shape, error

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	_, error := b.GetByIndex(i)

	if error != nil {
		return nil, error
	}

	result := make([]Shape, len(b.shapes))
	copy(result, b.shapes)
	b.shapes[i] = shape

	return result[i], nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum = sum + shape.CalcPerimeter()
	}

	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64

	for _, shape := range b.shapes {
		sum = sum + shape.CalcArea()
	}

	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	isExist := false

	for i := 0; i < len(b.shapes); i++ {
		_, ok := b.shapes[i].(*Circle)

		if ok {
			b.ExtractByIndex(i)
			isExist = true
			i--
		}
	}

	if !isExist {
		return errors.New("there are no circles in the list")

	}

	return nil

}
