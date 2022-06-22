func Smallest[T any](s []T) T {
	r := s[0]
	for _, v := range s[1:] {
		if v < r {
			r = v
		}
	}
	return r
}
