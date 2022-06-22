NewPair(1, int64(2))

func Double[E constraints.Integer](s []E) []E {
	r := make([]E, len(s))

	for i, v := range s {
		r[i] = v + v
	}

	return r
}
