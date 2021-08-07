package main

import (
	"fmt"
	"time"
)

func main() {
  fmt.Println("Inicio")

  c1 := make(chan chan int)
  //c2 := make(chan chan int)
  //c3 := make(chan chan int)
  x := make(chan int)
  //x <- 2
  //x <- 5

  go func() {
  	fmt.Println("Desde la primera")

  	c1 <- x
  	//c2 <- x
  	//c3 <- x
  	fmt.Printf("Canal 1 %v de tipo %T\n", c1, c1)
    //fmt.Printf("Canal 2 %v de tipo %T\n", c2, c2)
    //fmt.Printf("Canal 3 %v de tipo %T\n", c3, c3)
  }()
  go func() {
  	fmt.Println("Desde la segunda")

  	x <- 2
  	fmt.Println("Desde la segunda fin")
  }()

  go func() {
  	nc := <- c1
  	fmt.Printf("Esto es desde la tercera %v\n", nc)
  }()

  //<- x
  time.Sleep(4 * time.Second)
  //b := <- x
  //z := <- x
  //fmt.Println(b)
  //fmt.Println(z)
  //a := <- c1

  //fmt.Printf("AHhh %d\n", a)

}
