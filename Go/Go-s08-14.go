func Map[F, T any](s []F, f func(F) T) []T {
	// ...
}

var s []int

f := func(i int) int64 {
	return int64(i)
}

var r []int64

r = Map[int, int64](s, f)

r = Map[int](s, f)

r = Map(s, f)

Print[int]([]int{1, 2, 3})

s1 := []int{1, 2, 3}
Print(s1)

func Map[F, T any](s []F, f func(F) T) []T {
	r := make([]T, len(s))
	for i, v := range {
		r[i] = f(v)
	}
	return r
}

strs := Map([]int{1, 2, 3}, stconv.Itoa)

func NewPair[F any](f1, f2 F) *Pair[F] {
	// ...
}

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
	ID int
	Person struct {
		Name string
		Job  string
	}
}

const s = `{"ID": 42, "Person": {"Name": "George Constanza", "Job": "Architect"}}`
err := json.Unmarshal([]byte(s), &data)

if err != nil {
	log.Fatal(err)
}

fmt.Println(data.ID, data.Person.Name, data.Person.Job)

var s = []struct {
	i int
	s string
}{
	struct {
		i int
		s string
	}{6 * 9, "Question"},
	struct {
		i int
		s string
	}{42, "Answer"},
}

var t = []struct {
	i int
	s string
}{
	{6 * 9, "Question"},
	{42, "Answer"},
}
