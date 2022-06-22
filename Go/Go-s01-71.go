package main

type SyntaxError struct {
	msg    string // description of error
	Offset int64  // error occurred reading Offset bytes
}

func (e *SyntaxError) Error() string { return e.msg }

// func main() {
// 	if err := dec.Decode(&val); err != nil {
// 		if serr, ok := err.(*json.SyntaxError); ok {
// 			line, col := findLine(f, serr.Offset)
// 			return fmt.Errorf("%s:%d:%d: %v", f.Name(), line
// 				col, err)
// 		}
// 		return err
// 	}
// }
