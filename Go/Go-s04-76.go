func (r MyReader) Read(b []byte) (n int, err error) {
	b[0] = 'A'

	return 1, nil
}
