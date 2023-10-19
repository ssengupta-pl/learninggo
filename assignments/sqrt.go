package assignments

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot compute square root of -ve number: %f", float64(e))
}

func Sqrt(x float64) (float64, int, error) {
	if x < 0.0 {
		return 0.0, 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	iter := 1

	for {
		zpreviter := z
		z -= (z*z - x) / (2 * z)

		change := zpreviter - z

		if (iter > 1) && (change <= 0.000000000000001) {
			break
		} else {
			iter++
		}
	}

	return z, iter, nil
}

func Sqrt2(x float64) (float64, error) {
	if x < 0.0 {
		return 0.0, ErrNegativeSqrt(x)
	}

	z := 1.0
	iter := 1

	for {
		zpreviter := z
		z -= (z*z - x) / (2 * z)

		change := zpreviter - z

		if (iter > 1) && (change <= 0.000000000000001) {
			break
		} else {
			iter++
		}
	}

	return z, nil
}
