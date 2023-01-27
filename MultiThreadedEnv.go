import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d: Processing task %d\n", id, task)
		time.Sleep(time.Second)
		results <- task * 2
	}
}

func main() {
	tasks := make(chan int, 100)
	results := make(chan int, 100)

	for i := 0; i < 10; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 5; i++ {
		go worker(i, tasks, results)
	}

	for i := 0; i < 10; i++ {
		fmt.Print("Result", <-results)
	}
}