package main

import (
	"html/template"
	"log"
	"os"
)

type Foo struct {
	Bar string
}

func main() {
	tmpl := template.Must(template.New("home").Parse(`
<a title={{.Bar | html}}>
`))

	foo := Foo { "haha onclick=evil()" }

	if err := tmpl.Execute(os.Stdout, foo); err != nil {
		log.Fatal(err)
	}
}
