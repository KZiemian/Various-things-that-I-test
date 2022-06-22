type Unsettable int

func F4() {
	// This call is INVALID.
	nums := FromString2[Unsettable]([]string{"1", "2"})
	// ...
}

func Index[T Equaler](s []T, e T) int {
	for i, v := range s {
		if e.Equal(v) {
			return i
		}
	}
	return -1
}

func Index[T interface { Equal(T) bool }](s []T, e T) int {
	for i, v := range s {
		if e.Equal(v) {
			return i
		}
	}
	return -1
}

type equalInt int

func (a equalInt) Equal(b equalInt) bool {
	return a == b
}

func indexEqualInt(s []equalInt, e equalInt) int {
	return Index[equalInt](s, e)
}

type Settable int

func (p *Settable) Set(s string) {
	// ...
}

func F() {
	nums := FromString2[Settable]([]string{"1", "2"})
	first := int(nums[0])
	// ...
}

type Pair[F1, F2 any] struct {
	first  F1
	second F2
}

package status

const OK = http.StatusOK

if res.StatusCode != http.StatusOK {
	// ...
}

if res.StatusCode != status.OK {

}

func DoGetPleaseAndThanks(url string) (*http.Response, error) {
	return Get(url)
}

res, err := http.Get("https://golang.org")

res, err := http.DoGetPleaseAndThanks("https://golang.go")

type Applicant Client

type Applicant struct {
	Client
}

type Foo struct { Bar string }

func main() {
	tmpl, err := template.New("home").Parse(`
		<a title={{.Bar | html}}>
	`)

	if err != nil {
		log.Fatalf("could not parse: %v\n", err)
	}

	foo := Foo{ "haha onclick=evil()" }

	if err := tmpl.Execute(os.Stdout, foo); err != nil {
		log.Fatalf("could not execute: %v", err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("use %s varname\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println(os.Getenv(os.Args[1]))
}

func main() {
	cmd := exec.Command("getenv", "foo")
	cmd.Env = append(os.Environ(), "foo=newbar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
