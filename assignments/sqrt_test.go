package assignments

import (
	"math"
	"testing"
)

func TestSqrtWithPositiveNumber(t *testing.T) {
	result, iter, error := Sqrt(2.0)

	// If there is an error, then fail.
	if error != nil {
		t.Fatal(error)
	}

	// No error was thrown, check the result.
	expected := math.Sqrt(2.0)

	// If the difference is less than or equal to the tolerance,
	// we will consider the result to be the same as expected,
	// otherwise it is an error.
	tolerance := 0.0000001
	diff := math.Abs(expected - result)
	if diff > tolerance {
		t.Fatalf("Sqrt test failed: expected %f, received %f, diff %f", expected, result, diff)
	}

	// Finally, check the number of iterations.
	expectedIter := 6

	if expectedIter != iter {
		t.Fatalf("Sqrt test failed: expected iter %d, received %d", expectedIter, iter)
	}
}

func TestSqrtWithNegativeNumber(t *testing.T) {
	_, _, error := Sqrt(-2.0)

	// If there is no error, then fail.
	if error == nil {
		t.Fatal("Sqrt test failed: expected an error as the number was -ve")
	}

	// Error is thrown, make sure it says what is expected.
	expected := "Cannot compute square root of -ve number: -2.000000"
	received := string(error.Error())

	if expected != received {
		t.Fatalf("Sqrt test failed: error message mismatch expected %v, received %v", expected, received)
	}
}

func TestSqrt2WithPositiveNumber(t *testing.T) {
	result, error := Sqrt2(2.0)

	// If there is an error, then fail.
	if error != nil {
		t.Fatal(error)
	}

	// No error was thrown, check the result.
	expected := math.Sqrt(2.0)

	// If the difference is less than or equal to the tolerance,
	// we will consider the result to be the same as expected,
	// otherwise it is an error.
	tolerance := 0.0000001
	diff := math.Abs(expected - result)
	if diff > tolerance {
		t.Fatalf("Sqrt2 test failed: expected %f, received %f, diff %f", expected, result, diff)
	}
}

func TestSqrt2WithNegativeNumber(t *testing.T) {
	_, error := Sqrt2(-2.0)

	// If there is no error, then fail.
	if error == nil {
		t.Fatal("Sqrt2 test failed: expected an error as the number was -ve")
	}

	// Error is thrown, make sure it says what is expected.
	expected := "Cannot compute square root of -ve number: -2.000000"
	received := string(error.Error())

	if expected != received {
		t.Fatalf("Sqrt2 test failed: error message mismatch expected %v, received %v", expected, received)
	}
}
