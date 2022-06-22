type Func struct {
	in  reflect.Type
	out reflect.Type
	f func(interface{}) interface{}
}

func (f Func) Call(v interface{}) interface{} {
	return f.f(v)
}

func NewFunc(f interface{}) (*Func, error) {
	return &Func{
		in: t
		f: func(x interface{}) interface{} {
			out := vf.Call([]reflect.Value{reflect.ValueOf(x)})
			return out[0].Interface{}
		},
	}, nil
}

func Must(f *Func, err error) *Func {
	if err != nil {
		panic(err)
	}
	return f
}

func Map(f *Func, vs interface{}) interface{}

type List struct {
	Head interface{}
	Tail *List
}

func Map(f *Func, l *List) *List {
	if l == nil {
		return nil
	}
	return &List{f.Call(l.Head), Map(f, l.Tail)}
}

func (l *List) Map(f *Func) *List {
	if l == nil {
		return nil
	}

	return &List{
		f.Call(l.Head)
		Map(f, l.Tail)
	}
}

toUpper := Must(NewFunc(strings.ToUpper))

m := &List{"Hello", &List{"bye", nil}}
res := m.Map(toUpper)
fmt.Println(res)

func (l *List) Map(f *Func) *List

type Mapper interface {
	Map(*Func) ???
}

type Maybe struct {
	Value interface{}
}

func (m Maybe) Map(f *Func) Maybe {
	if m.Value == nil {
		return Maybe{}
	}
	return Maybe{
		f.Call(m.Value)
	}
}

toUpper := Must(NewFunc(strings.ToUpper))

m := Maybe{}
res := m.Map(toUpper)
fmt.Println(res.Value)

toUpper := Must(NewFunc(strings.ToUpper))
twice := Must(NewFunc(func(s string) string { return s + s }))

m := Maybe{"hello"}
res := m.Map(toUpper).Map(twice)
fmt.Println(res.Value)

type Person struct {
	address *Address
}

type Address struct {
	city *City
}

type City struct {
	weather *Weather
}

type Weather struct {
	desc string
}

func (p Person) Address() *Address {
	return p.address
}

func (a Address) City() *City {
	return a.city
}

func (c City) Weather() *Weather {
	return c.weather
}

func (w Weather) Desc() string {
	return w.desc
}

func (p Person) Weather() string {
	a := p.Address()
	if a == nil {
		return "no weather"
	}
	c := a.City()
	if c == nil {
		return "no weather"
	}
	w := c.Weather()
	if w == nil {
		return "no weather"
	}
	return w.Desc()
}

func (p Person) Weather() string {
	m := Maybe{p}.
		Map(Must(NewFunc(func(a Address) *City {
			return a.city
		})))
}

func (p Person) Weather() string {
	w := Maybe{p}.
		Map(Must(NewFunc(Person.Address))).
		Map(Must(NewFunc(Address.City))).
		Map(Must(NewFunc(City.Weather))).
		Map(Must(NewFunc(Weather.Description)))
	if w.Value == nil {
		return "no weather"
	}
	return w.Value.(string)
}

func (m Maybe) Do(fs ...interface{}) (Maybe, error) {
	if len(fs) == 0 {
		return m, nil
	}
	f, err := NewFunc(fs[0])
	if err != nil {
		return Maybe{}, err
	}
	return m.Map(f).Do(fs[1:]...)
}

func (p Person) Weather() string {
	w := Maybe{p}.Do(
		Person.Address,
		Address.City,
		City.Weather,
		Weather.Desc
	)

	if w.Value == nil {
		return "no weather"
	}
	return w.Value.(string)
}
