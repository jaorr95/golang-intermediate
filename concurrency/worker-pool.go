package main

import "fmt"

func main() {

	tasks := []int{1,5,3,8,40,10}
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))
	nWorkers := 3
	initWorkers(nWorkers, jobs , results)

	for _, task := range tasks {
		jobs <- task
	}

	close(jobs)

	for i := 0; i < len(tasks); i++ {
		<- results
	}


}

func initWorkers(quantity int, jobs chan int, results chan int) {
	for i := 0; i < quantity; i++ {
		go Worker(i, jobs, results)
	}
}


func Worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("El worker %d esta ejecetuando el job %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("El worker %d proceso el job %d con resultado %d\n",id, job, fib)
		result <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}