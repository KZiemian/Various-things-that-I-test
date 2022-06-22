var m map[string]int

m = make(map[string]int)

type Node struct {
	Next  *Node
	Value interface{}
}

var first *Node

visited := make(map[*Node]bool)

for n := first; n != nil; n = n.Next {
	if visitied[n] {
		fmt.Println("cycle detected")
		break
	}
	visited[n] = true
	fmt.Println(n.Value)
}

type PrimaryColour int

const (
	Red = iota
	Yellow
	Blue
)

func (c PrimaryColour) String() string {
	switch c {
	case Red:
		return "Red"

	case Yellow:
		return "Yellow"

	case Blue:
		return "Blue"

	default:
		panic(fmt.Sprintf("unknown colour %d", int(c)))
	}
}

type MyLongTypeName struct {
	...
}

func foo(x int, y string) (MyLongTypeName, error) {
	if x == 42 {
		return MyLongTypeName{}, error.New("x cannot be the meaning of life")
	}

	if y == "" {
		return MyLongTypeName{}, error.New("y cannot be empty")
	}

	return MyLongTypeName{...}, nil
}


func GetKeyURL(x int) (*url.URL, error) {
	y := Foo(x)?
	z := Bar(y)?
	return Baz(z)
}

func Foo(x int) (int, error) {
	...
}

func Bar(y string) (string, error) {
	...
}

func Baz(z string) (*url.URL, error) {
	...
}

type Person struct {
	Name  string
	Likes []string
}

var people []*Person

likes := make(map[string][]*Person)

for _, p := range people {
	for _, l := range p.Likes {
		likes[l] = append(likes[l], p)
	}
}

for _, p := range likes["cheese"] {
	fmt.Println(p.Name, "likes cheese.")
}

fmt.Println(len(likes["bacon"]), "people like bacon.")
