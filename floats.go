package grbo

import (
	"math/rand"
)

func float64sSub(a, b []float64) []float64 {
	ret := make([]float64, len(a))
	for i, v := range a {
		ret[i] = v - b[i]
	}
	return ret
}

func float64sAdd(a, b []float64) []float64 {
	ret := make([]float64, len(a))
	for i, v := range a {
		ret[i] = v + b[i]
	}
	return ret
}

func float64sProduct(a, b []float64) []float64 {
	ret := make([]float64, len(a))
	for i, v := range a {
		ret[i] = v * b[i]
	}
	return ret
}

func float64sRandom(n int, lower, upper float64) []float64 {
	ret := make([]float64, n)
	for i := 0; i < n; i++ {
		ret[i] = lower + rand.Float64() * (upper - lower)
	}
	return ret
}

func floats(v float64, n int) []float64 {
	ret := make([]float64, n)
	for i, _ := range ret {
		ret[i] = v
	}
	return ret
}