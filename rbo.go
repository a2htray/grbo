package grbo

import (
	"math"
	"math/rand"
)

type RBO struct {
	t int // 迭代次数
	m int // 解的个数
	n int // 解的维数
	c int // 查找次数
	objectFunc func([]float64) float64 // 目标函数
	lowerLimit []float64
	upperLimit []float64
	historyBests []*Rice
	rices []*Rice
}

func (r *RBO) HistoryBest() []*Rice {
	return r.historyBests
}

func (r *RBO) T() int {
	return r.t
}

func (r *RBO) M() int {
	return r.m
}

func (r *RBO) N() int {
	return r.n
}

func (r *RBO) Run() {
	size := r.m / 3
	for t := 0; t < r.t; t++ {
		sortedInds := r.argsort()
		maintainerInds, restorerInds, sterileInds := sortedInds[:size], sortedInds[size:2*size], sortedInds[2*size:]
		// Hybridization
		for _, sterileInd := range sterileInds {
			maintainer := r.rices[intsChoice(maintainerInds)]
			copySterile := copyFrom(r.rices[sterileInd])
			for j := 0; j < r.n; j++ {
				r1, r2 := rand.Float64(), rand.Float64()
				oldGene := copySterile.genes[j]
				copySterile.genes[j] = (r1 * copySterile.genes[j] + r2 * maintainer.genes[j]) / (r1 + r2)
				newFitness := r.objectFunc(copySterile.genes)
				if newFitness < copySterile.fitness {
					copySterile.fitness = newFitness
				} else {
					copySterile.genes[j] = oldGene
				}
			}
			r.rices[sterileInd] = copySterile
		}
		// Selfing
		best := r.rices[maintainerInds[0]]
		for _, restorerInd := range restorerInds {
			// Renewal
			if r.rices[restorerInd].c >= r.c {
				r.rices[restorerInd] = &Rice{
					c:       0,
					genes:   genes(r.n, r.lowerLimit, r.upperLimit),
					fitness: 0,
				}
				break
			}
			randomRestorerInd := intsChoice(restorerInds)
			for restorerInd == randomRestorerInd {
				randomRestorerInd = intsChoice(restorerInds)
			}

			copyRestorer := copyFrom(r.rices[restorerInd])
			randomRestorer := r.rices[randomRestorerInd]
			for j := 0; j < r.n; j++ {
				r3 := rand.Float64()
				oldGene := copyRestorer.genes[j]
				copyRestorer.genes[j] = r3 * (best.genes[j] - randomRestorer.genes[j]) + copyRestorer.genes[j]
				newFitness := r.objectFunc(copyRestorer.genes)
				if newFitness < copyRestorer.fitness {
					copyRestorer.c++
					copyRestorer.fitness = newFitness
				} else {
					copyRestorer.genes[j] = oldGene
				}
			}
		}


		r.historyBests[t] = copyFrom(r.rices[r.argsort()[0]])
	}
}

func (r *RBO) argsort() []int {
	fitnessList := make([]float64, r.m)
	for i := 0; i < r.m; i++ {
		fitnessList[i] = r.objectFunc(r.rices[i].Values())
		r.rices[i].fitness = fitnessList[i]
	}

	return argsortFunc(fitnessList)
}

func New(m, n int, options ...Option) *RBO {
	rbo := &RBO{
		t: 100,
		m: m,
		n: n,
		c: 3,
		objectFunc: func(fs []float64) float64 {
			var fitness float64
			for _, f := range fs {
				fitness = math.Pow(f, 2)
			}
			return fitness
		},
		lowerLimit: floats(-1, n),
		upperLimit: floats(1, n),
	}
	for _, option := range options {
		option(rbo)
	}

	rbo.historyBests = make([]*Rice, rbo.t)
	rbo.rices = rices(rbo.m, rbo.n, rbo.c, rbo.lowerLimit, rbo.upperLimit)

	return rbo
}


