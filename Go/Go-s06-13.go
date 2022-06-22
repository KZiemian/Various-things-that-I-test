package main

import "fmt"

func main() {
	// ...
	for _, task := range tasks {
		taskCh <- task
	}
}

func worker(taskCh chan Task) {
	for {
		task := <-taskCh
		process(task)
	}
}
