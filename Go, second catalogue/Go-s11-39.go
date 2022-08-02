reader := bufio.NewReader(os.Stdin)
char, _, err := reader.ReadRune()

if err != nil {
	fmt.Println(err)
}

fmt.Println(char)

switch char {
case 'A':
	fmt.Println("A Key Pressed")

case 'a':
	fmt.Println("a Key Pressed")
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 10; i++ {
		scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}