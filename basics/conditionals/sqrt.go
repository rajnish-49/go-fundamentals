package main

import (
	"math"
)


func Sqrt(x float64) float64 {
	z := 1.0

	for {
		next := z - (z*z - x) / (2*z)

		// stop when change is very small
		if math.Abs(next-z) < 1e-6 {
			return next
		}

		z = next
	}
}