package grbo

import "math/rand"

func intsChoice(ints []int) int {
	n := len(ints)
	return ints[rand.Intn(n)]
}