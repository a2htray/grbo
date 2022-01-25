package grbo

type Rice struct {
	c int
	genes []float64
	fitness float64
}

func (r *Rice) Values() []float64 {
	return r.genes
}

func (r *Rice) Fitness() float64 {
	return r.fitness
}

func (r *Rice) At(i int) float64 {
	return r.genes[i]
}

func rices(m, n, c int, lowerLimit, upperLimit []float64) []*Rice {
	rs := make([]*Rice, m)
	for i := 0; i < m; i++ {
		rs[i] = &Rice{
			c:       c,
			genes:   genes(n, lowerLimit, upperLimit),
			fitness: 0,
		}
	}
	return rs
}

func genes(n int, lowerLimit, upperLimit []float64) []float64 {
	return float64sAdd(lowerLimit, float64sProduct(float64sSub(upperLimit, lowerLimit), float64sRandom(n, 0, 1)))
}

func copyFrom(rice *Rice) *Rice {
	copyRice := &Rice{
		c:       rice.c,
		genes:   nil,
		fitness: rice.fitness,
	}
	copyGenes := make([]float64, len(rice.genes))
	for i, gene := range rice.genes {
		copyGenes[i] = gene
	}
	copyRice.genes = copyGenes
	return copyRice
}
