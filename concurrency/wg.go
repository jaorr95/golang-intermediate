package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go goroutine(i, &wg)
	}
	wg.Wait()

}

func goroutine(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Iniciando %d..\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Desperteee %d..\n", i)

}
