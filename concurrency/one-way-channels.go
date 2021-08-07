package main

import "fmt"

func main() {
  genChan := make(chan int)
  doubleChan := make(chan int)

  go generator(genChan)
  go double(genChan, doubleChan)
  print(doubleChan)
}


func generator(c chan<- int){

	for i := 1; i<= 10; i++ {
		c <- i
	}
    //necesario hacer close para terminar el loop de double
	close(c)

}

func double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value
	}
    //necesario hacer el close para terminar loop de print
	close(out)
}

func print(c <-chan int) {

	for value := range c {
		fmt.Println(value)
	}

}
