type X interface {
	X()
}

type NotX interface {
	X() int
}

type S interface {
	A | B
}

type X interface { X() }
type NotX interface { X() int }
type Y interface { Y() int }
type NotY interface { Y() int }
type Z interface { Z() }
type NotZ interface { Z() int }

type S interface {
	NotX | Z
	Y | Z
	X | NotY
}

type TertiumNonDatur interface {
	X | NotX
	Y | NotY
	Z | NotZ
}

type S interface {
	TertiumNonDatur

	NotX | Z
	Y | Z
	X | NotY
}
