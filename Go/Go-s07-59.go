data := struct {
	Title               string
	Firstname, Lastname string
	Rank                int
}{
	"Dr", "Carl", "Sagan", 7,
}

if err := tmpl.Execute(os.Stdout, data); err != nil {
	log.Fatal(err)
}

b, err := json.Marshal(struct {
	ID   int
	Name string
}{42, "The answer"})

if err != nil {
	log.Fatal(err)
}
fmt.Printf("%s\n", b)

var data struct {
	ID   int
	Name string
}
err := json.Unmarshal([]byte(`{"ID": 42, "Name": "The answer"}`), &data)
if err != nil {
	log.Fatal(err)
}
fmt.Println(data.ID, data.Name)


var data struct {
	ID     int
	Person struct {
		Name string
		Job  string
	}
}

const s = `{"ID":42,"Person":{"Name":"George Costanza","Jab":"Architect"}}`

err := json.Unmarshal([]byte(s), &data)
if err != nil {
	log.Fatal(err)
}
fmt.Println(data.ID, data.Person.Name, data.Person.Job)

func TestIndex(t *testing.T) {
	var tests = []struct {
		s   string
		sep string
		out int
	}{
		{"", "", 0},
		{"", "a", -1},
		{"fo", "foo", -1},
		{"foo", "foo", 0},
		{"oofofoofooo", "f", 2},
	}
	for _, test := range tests {
		actual := strings.Index(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Index(%q, %q) = %v; want %v", test.s,
				test.sep, acutal, test.out)
		}
	}
}

var (
	viewCount   int64
	viewCountMu sync.Mutex
)

var viewCount struct {
	sync.Mutex
	n int64
}

viewCount.Lock()
viewCount.n++
viewCount.Unlock()
