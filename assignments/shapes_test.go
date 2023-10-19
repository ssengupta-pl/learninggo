package assignments

import (
	"math"
	"testing"
)

func TestCircleAreaWithPositiveRadius(t *testing.T) {
	circle := Circle{1.0}

	result, error := circle.Area()

	// Check for errors.
	if error != nil {
		t.Fatal(error)
	}

	// Now check results.
	expected := math.Pi * 1.0 * 1.0

	if result != expected {
		t.Fatalf("Test circle area with +ve r failed: expected - %f, received - %f", expected, result)
	}
}

func TestCircleAreaWithNegativeRadius(t *testing.T) {
	circle := Circle{-1.0}

	_, error := circle.Area()

	// Check if errors exist - there should be one at least.
	if error == nil {
		t.Fatal(error)
	}

	// Error exists, now check if that is what you expect.
	expected := "Invalid circle: radius -1.000000"
	received := string(error.Error())

	if received != expected {
		t.Fatalf("Test circle area with -ve r failed: expected - %v, received - %v", expected, received)
	}
}

func TestRectangleWithPositiveDims(t *testing.T) {
	rectangle := Rectangle{2.0, 1.0}

	result, error := rectangle.Area()

	// Check for errors.
	if error != nil {
		t.Fatal(error)
	}

	// Now check results.
	expected := 2.0 * 1.0

	if result != expected {
		t.Fatalf("Test rectangle area with +ve dims failed: expected - %f, received - %f", expected, result)
	}
}

func TestRectangleWithNegativeLength(t *testing.T) {
	rectangle := Rectangle{-2.0, 1.0}

	_, error := rectangle.Area()

	// Check if errors exist - there should be one at least.
	if error == nil {
		t.Fatal(error)
	}

	// Error exists, now check if that is what you expect.
	expected := "Invalid rectangle: length -2.000000, width 1.000000"
	received := string(error.Error())

	if received != expected {
		t.Fatalf("Test rectangle area with -ve length failed: expected - %v, received - %v", expected, received)
	}
}

func TestRectangleWithNegativeWidth(t *testing.T) {
	rectangle := Rectangle{2.0, -1.0}

	_, error := rectangle.Area()

	// Check if errors exist - there should be one at least.
	if error == nil {
		t.Fatal(error)
	}

	// Error exists, now check if that is what you expect.
	expected := "Invalid rectangle: length 2.000000, width -1.000000"
	received := string(error.Error())

	if received != expected {
		t.Fatalf("Test rectangle area with -ve width failed: expected - %v, received - %v", expected, received)
	}
}

func TestSquareWithPositiveSide(t *testing.T) {
	square := Square{1.0}

	result, error := square.Area()

	// Check for errors.
	if error != nil {
		t.Fatal(error)
	}

	// Now check results.
	expected := 1.0 * 1.0

	if result != expected {
		t.Fatalf("Test square area with +ve side failed: expected - %f, received - %f", expected, result)
	}

}

func TestSquareWithNegativeSide(t *testing.T) {
	square := Square{-1.0}

	_, error := square.Area()

	// Check if errors exist - there should be one at least.
	if error == nil {
		t.Fatal(error)
	}

	// Error exists, now check if that is what you expect.
	expected := "Invalid square: side -1.000000"
	received := string(error.Error())

	if received != expected {
		t.Fatalf("Test square area with -ve side failed: expected - %v, received - %v", expected, received)
	}
}
