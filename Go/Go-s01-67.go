// package main

// import "fmt"

// func Open(name string) (file *File, err error)

// type error interface {
// 	Error() string
// }

// errorString is a trivial implementation of error.
// type errorString struct {
// 	s string
// }

// func (e *errorString) Error() string {
// 	return e.s
// }

// New returns an error that formats as the given text.
// func New(text string) error {
// 	return &errorString{text}
// }

// func main() {
// 	f, err := os.Open("filename.ext")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// do something with the open *File f
// }
