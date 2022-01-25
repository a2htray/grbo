package grbo

import (
	"math"
	"testing"
)

func TestRBO_Run(t *testing.T) {
	rbo := New(12, 3)
	rbo.Run()
	for i, rice := range rbo.HistoryBest() {
		t.Logf("iter %d %v %v", i, rice.Values(), rice.Fitness())
	}
}

func TestRBO_Run2(t *testing.T) {
	rbo := New(60, 10, WithT(3000), WithC(50), WithLowerLimit([]float64{
		-100, -100, -100, -100, -100, -100, -100, -100, -100, -100,
	}), WithUpperLimit([]float64{
		100, 100, 100, 100, 100, 100, 100, 100, 100, 100,
	}), WithObjectFunc(func(fs []float64) float64 {
		part := 0.0
		for i := 1; i < len(fs); i++ {
			part += math.Pow(fs[i], 2)
		}
		return math.Pow(fs[0], 2) + math.Pow(10, 6) * part
	}))
	rbo.Run()
	for i, rice := range rbo.HistoryBest() {
		t.Logf("iter %d %v %v", i, rice.Values(), rice.Fitness())
	}
}