type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

func DotProduct[T Numeric](s1, s2 []T) T {
	if len(s1) != len(s2) {
		panic("DotProduct: slice of unequal length")
	}
	var r T
	for i := range s1 {
		r += s1[i] * s2[i]
	}
	return r
}

type NumericAbs[T any] interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128
	Abs() T
}

func AbsDifference[T NumericAbs](a, b T) T {
	d := a - b
	return d.Abs()
}

type OrderedNumeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type OrderedAbs[T OrderedNumeric] T

func (a OrderedAbs[T]) Abs() OrderedAbs[T] {
	if a < 0 {
		return -a
	}
	return a
}

type ComplexAbs[T Complex] T

func (a ComplexAbs[T]) Abs() ComplexAbs[T] {
	d := math.Hypot(float64(real(a)), float64(imag(a)))
	return ComplexAbs[T](complex(d, 0))
}

func OrderedAbsDifference[T OrderedNumeric](a, b T) T {
	return T(AbsDifference(OrderedAbs[T](a), OrderedAbs[T](b)))
}

func ComplexAbsDifference[T Complex](a, b T) T {
	return T(AbsDifference(ComplexAbs[T](a), ComplexAbs[T](b)))
}

// INVALID

func GeneralAbsDifference[T Numeric](a, b T) T {
	switch (interface{})(a).(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64:
		return OrderedAbsDifference(a, b) // INVALID
	case complex64, complex128:
		return CompleAbsDifference(a, b) // INVALID
	}
}

type VectorAlias = Vector

var v VectorAlias[int]

type VectorInt = Vector[int]

var PrintInts = Print[int]

type Lockable[T any] struct {
	T
	mu sync.Mutex
}

func (l *Lockable[T]) Get() T {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.T
}

func (l *Lockable[T]) Set(v T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.T = v
}
