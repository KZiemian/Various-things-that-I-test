// package main

func main() {
	tasks := getTasks()

	for _, task := range tasks {
		process(task)
	}
	// ...
}

func main() {
	taskCh := make(chan Task, 3)

	for i := 0; i < numWorkers; i++ {
		go worker(ch)
	}

	hellaTasks := getTasks()

	for _, task := range hellaTasks {
		taskCh <- task
	}
}
