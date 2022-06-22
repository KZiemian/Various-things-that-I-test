if err != nil {
	return err
}

scanner := bufio.NewScanner(input)
for scanner.Scan() {
	token := scanner.Text()
	// process token
}
if err := scanner.Err(); err != nil {
	// process the error
}

func (s *Scanner) Scan() (token []byte, error)

scanner := bufio.NewScanner(input)
for {
	token, err := scanner.Scan()
	if err != nil {
		return err // or maybe break
	}
	// process token
}

_, err = fd.Write(p0[a:b])
if err != nil {
	return err
}
_, err = fd.Write(p1[c:d])
if err != nil {
	return err
}
_, err = fd.Write(p2[e:f])
if err != nil {
	return err
}
// and so on

var err error
write := func(buf []byte) {
	if err != nil {
		return
	}
	_, err = w.Write(buf)
}
write(p0[a:b])
write(p1[c:d])
write(p2[e:f])
// and so on
if err != nil {
	return err
}

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) write(buf []byte) {
	if ew.err != nil {
		return
	}
	_, ew.err = ew.w.Write(buf)
}

ew := &errWriter{w: fd}
ew.write(p0[a:b])
ew.write(p1[c:d])
ew.write(p2[e:f])
// and so on
if ew.err != nil {
	return ew.err
}

b := bufio.NewWriter(fd)
b.Write(p0[a:b])
b.Write(p1[c:d])
b.Write(p2[e:f])
// and so on
if b.Flush() != nil {
	return b.Flush()
}
