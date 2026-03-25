package main

import (
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {  // v is only available in the if block
		return v
	}
	return lim
}