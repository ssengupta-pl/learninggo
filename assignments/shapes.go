package assignments

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() (float64, error)
}

type ShapeDimensionsError string

func (e ShapeDimensionsError) Error() string {
	return string(e)
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Square struct {
	side float64
}

func (circle Circle) Area() (float64, error) {
	r := circle.radius

	if r < 0 {
		return 0.0, ShapeDimensionsError(fmt.Sprintf("Invalid circle: radius %f", r))
	}

	return math.Pi * r * r, nil
}

func (rectangle Rectangle) Area() (float64, error) {
	l := rectangle.length
	w := rectangle.width

	if l < 0 || w < 0 {
		return 0.0, ShapeDimensionsError(fmt.Sprintf("Invalid rectangle: length %f, width %f", l, w))
	}

	return l * w, nil
}

func (square Square) Area() (float64, error) {
	s := square.side

	if s < 0 {
		return 0.0, ShapeDimensionsError(fmt.Sprintf("Invalid square: side %f", s))
	}

	return s * s, nil
}
