buf := make([]byte, 100)
n, err := r.Read(buf)
buf = buf[:n]
if err == io.EOF {
	log.Fatal("read failed:", err)
}

func readfile(path string) error {
	err := openfile(path)

	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
}

func main() {
	err := readfile(".bashrc")

	if strings.Contains(error.Error(), "not found") {
		//
	}
}
