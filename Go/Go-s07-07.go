// package main

// import "fmt"

type Many struct {
	Head interface{}
	Tail *Many
}

func (m *Many) Map(f *Func) *Many {
	if m == nil {
		return nil
	}

	res := m.Tail.Map(f)

	for _, v := range toSlice(f.Call(m.Head)) {
		r = &Many{v, res}
	}
	return res
}

toUpper := Must(NewFunc(strings.ToUpper))

m := Many{"hello there", "good bye"}
res := m.Map(toUpper)
fmt.Println(res)

toUpper := Must(NewFunc(strings.ToUpper))
fields := Must(NewFunc(strings.Fields))

m := Many{"hello there", "good bye"}
res := m.Map(toUpper).Map(fields)
fmt.Println(res.Value)

type Library struct {
	books []Book
}

type Book struct {
	pages []Page
}

type Page struct {
	lines []Line
}

type Line struct {
	text string
}

func (l Library) Books() []Book {
	return l.books
}

func (b Book) Pages() []Page {
	return b.pages
}

func (p Page) Lines() []Line {
	return p.lines
}

func (l Line) Text() string {
	return l.text
}

words := make(map[string]int)

for _, b := range l.Books() {
	for _, p := range b.Pages() {
		for _, l := range p.Lines() {
			for _, word := range strings.Fields(l.Text()) {
				words[word]++
			}
		}
	}
}

NewMany(l).
	Map(Must(NewFunc(Library.Books))).
	Map(Must(NewFunc(Book.Pages))).
	Map(Must(NewFunc(Pages.Lines))).
	Map(Must(NewFunc(Line.Text))).
	Map(Must(NewFunc(strings.Fields))).
	Map(Must(NewFunc(func (s string) bool {
		count[s]++;
		return true
	})))

_, err := NewMany(m).Do(
	Library.Books,
	Book.Pages,
	Page.Lines,
	Line.Text,
	strings.Field,
	func(s string) bool {
		count[s]++
		return true})

if err != nil {
	...
}

w, err := NewMany(m).
	Do(
		Library.Books,
		Book.Pages,
		Page.Lines,
		Line.Text,
		strings.Fields)

if err != nil {
	...
}

w.Each(func(s string) {
	count[s]++
})
