package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSleep(d1, c1, 1)
	go DoSleep(d2, c2, 2)

	//fmt.Println(<-c1)
	//fmt.Println(<-c2)

	for i := 0; i < 2 ; i++ {
		select {
		case channel1 := <- c1:
			fmt.Println(channel1)
		case channel2 := <- c2:
			fmt.Println(channel2)

		}
	}

}

func DoSleep(t time.Duration, c chan<- int, id int) {
	time.Sleep(t)
	c <- id

}
