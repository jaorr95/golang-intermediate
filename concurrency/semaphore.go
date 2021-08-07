package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Esta implemantacion de buffered channel liberando en la goroutina
	// permite controlar la cantidad de goroutinas que pueden estar ejecutandose al mismo tiempo
	// La cantidad del bufer limita la cantidad de goroutinas que se ejecutan al mismo tiempo
	c := make(chan int, 5)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		c <- 1
		go goroutineLimited(i, &wg, c)
	}
	wg.Wait()

}


func goroutineLimited(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Iniciando %d...\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Derpertando %d...\n", i)
	<-c
}
