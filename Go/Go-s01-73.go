package main

func main() {
	if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
		time.Sleep(1e9)
		continue
	}
	if err != nil {
		log.Fatal(err)
	}
}
