type ReadCloser interface {
	Read(b []byte) (n int, err os.Error)
	Close()
}

func ReadAndClose(r ReadCloser, buf []bytes) (n int, err os.Error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	r.Close()
	return
}


type Stringer interface {
	String() string
}

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.Ftoa(v, 'g', -1)
	}
	return "???"
}

type Binary uint64

func (i Binary) String() string {
	return strconv.Uitob64(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

var any interface
s := any.(Stringer)
for i := 0; i < 100; i++ {
	fmt.Println(s.String())
}
