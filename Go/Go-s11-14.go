type MyString string

type AnyString interface{ ~string }

type AnyString interface {
	~string
}

type ApproximateMyString interface {
	~MyString // INVALID: undelying type of MyString is not MyString
}

type ApproximateParameter[T any] interface {
	~T // INVALID: T is a type parameter
}

int | float32

~int8 | ~int16 | ~int32 | ~int64

type PredeclaredSignedInteger interface {
	int | int8 | int16 | int32 | int64
}

type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

package constraints

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Smallest[T constraints.Ordered](s []T) T {
	r := s[0] // panics if slice is empty

	for _, v := range s[1:] {
		if v < r {
			r = v
		}
	}

	return r
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}

	return -1
}

type ComparableHasher interface {
	comparable
	Hash() uintptr
}

type ImpossibleConstraint interface {
	comparable
	[]int
}

package graph

type NodeConstraint[Edge any] interface {
	Edges() []Edge
}

type EdgeConstraint[Node any] interface {
	Nodes() (from, to Node)
}

type Graph[Node NodeConstraint[Edge], EdgeConstraint[Node]] struct {
	// ...
}

func New[Node NodeConstraint[Edge], Edge EdgeConstraint[Node]](nodes []Node) *Graph[Node, Edge] {
	// ...
}

type ID string

func Call(ctx context.Context, id ID, req Message) (Message, error)

r, err := framework.Call(ctx, somepkg.ID, somepkg.Request{...})
resp := r.(somepkg.Response)

type ID string
func Call[Req Message, Resp Message](ctx contex.Context, id ID, req Req) (Resp, error)

resp, err := framework.Call[somepkg.Request, somepkg.Response](ctx,
	somepkg.ID, somepkg.Request{...})

type ID[Req Message, Resp Message] string
func Call[Req Message, Resp Message](ctx contex.Context, id ID[Req, Resp],
	req Req) (Resp, error)

var ID = framework.ID[MyRequest, MyResponse]("my-name")

resp, err := framework.Call(ctx, somepkg.ID, somepkg.Request)
