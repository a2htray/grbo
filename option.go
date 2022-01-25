package grbo

type Option func(r *RBO)

func WithT(t int) Option {
	return func(r *RBO) {
		r.t = t
	}
}

func WithC(c int) Option {
	return func(r *RBO) {
		r.c = c
	}
}

func WithLowerLimit(limit []float64) Option {
	return func(r *RBO) {
		r.lowerLimit = limit
	}
}

func WithUpperLimit(limit []float64) Option {
	return func(r *RBO) {
		r.upperLimit = limit
	}
}

func WithObjectFunc(objectFunc func([]float64) float64) Option {
	return func(r *RBO) {
		r.objectFunc = objectFunc
	}
}
