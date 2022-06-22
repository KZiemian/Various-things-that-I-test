hits := make(map[string]map[string]int)

n := hits["/doc/"]["au"]

func add(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}

add(hits, "/doc/", "au")

type Key struct {
	Path, Country string
}

hits := make(make[Key]int)

hits[Key{"/", "vn"}]++

n := hits[Key{"/ref/spec", "ch"}]

var counter = struct{
	sync.RWMutex
	m map[string]int
}{m: make(map[string]int)}

counter.RLock()
n := counter.m["some_key"]
counter.RUnlock()
fmt.Println("some_key", n)

counter.Lock()
counter.m["some_key"]++
counter.Unlock()

func (d *Document) Save(rwc io.ReadWriteCloser) error

func (d *Document) Save(wc io.WriteCloser) error

type NopCloser struct {
	io.Writer
}

// Close has no effect on the underlying writer.
func (c *NopCloser) Close() error {
	return nil
}

func (d *Document) Save(w io.Writer) error

func (d *Document) WriteTo(w io.Writer) error

func ReadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// ...
}

func comp(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func comp(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func comp(a, b int) int {
	if a > b {
		a = b
	} else if a < b {
		return 1
	} else {
		return 0
	}
}

func comp(a, b int) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
