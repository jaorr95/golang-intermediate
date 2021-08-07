package main

import "fmt"

func main() {
 //unbufferedChannelError()
 unbufferedChannelSuccess()
 bufferedChannel()
}

func unbufferedChannelError() {
	// Da error porque se bloquea el hilo al no tener espacio el channel
	//  Esto solo es posible si se ejecuta en hilos separados
	c := make(chan int)

	c <- 1

	fmt.Println(<-c)

}

func unbufferedChannelSuccess()  {
	c := make(chan int)

	go func() {
		fmt.Println("Otra goroutine")
		c <- 1
	}()

	fmt.Println(<-c)
}

func bufferedChannel() {
    // En este caso no se bloquea el hilo principal debido a que es channel tiene un buffer
	// si el bufer se llena pasaria lo mismo de unbufferedChannelError
	c := make(chan int, 3)

	c <- 2
	c <- 2
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
