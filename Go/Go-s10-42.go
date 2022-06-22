func Nextafter32(x, y float32) float32

func Pow(x, y float64) float64

func Pow10(n int) float64

func Remainder(x, y float64) float64

func Round(x float64) float64

func RountToEven(x floa64) float64

func Signbit(x float64) bool

func Sin(x float64) float64

func Sincos(x float64) (sin, cos float64)

func Sinh(x float64) float64

func Sqrt(x float64) float64

func Tan(x float64) float64

func Tanh(x float64) float64

func Y0(x float64) float64

func Y1(x float64) float64

func Yn(n int, x float64) float64

var matrix [,]float64

matrix = make([,]float64, 15, 11)

m.At(i, j)

m.AtSet(i, j, x)

c.AtSet(i, j, a.At(i, k) * b.At(k, ind.At(j)))

c[i, j] = a[i, k] * b[k, ind[j]]

a[i, j] => a.At(i, j)

x[i] => x.[](i)

x[i, j] => x.[](i, j)

x[i, j, k] = y => x.[]=(i, j, k, y)

x + y => x.+(y)

x[i, j] => x.AT__(i, j)

x[i, j] = y => x.ATSET__(i, j, y)

x + y => x.ADD__(y)

type Point struct {
	X, Y int
}

func (a Point) +(b Point) Point {
	return Point{...}
}

var a, b, c Point

c := a + b

type Point struct {
	X, Y int
}

func (a Point) ADD__(b Point) Point {
	return Point{...}
}

var a, b, c Point

c := a.ADD__(b)

type Matrix struct {
	array       []float64
	len, stride [2]int
}

func NewMatrix(n, m int) *Matrix
func (m *Matrix) []  (i, j int) float64
func (m *Matrix) []= (i, j int, x float64)

for i := 0; i < n; i++ {
	for j := 0; j < p; j++ {
		var t float64
		for k := 0; k < m; k++ {
			t += a[i, k] *
				b[k, j]
		}
		c[i, j] = t
	}
}
