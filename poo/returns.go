package main

import "fmt"

func sum(values ...int) int{
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func printNames(names ...string) {
	for _,value := range names {
		fmt.Println(value)
	}
}

func namedReturn(x int) (double int, triple int, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4
	return
}


func main() {
  fmt.Println(sum(1))
  fmt.Println(sum(1, 2))
  fmt.Println(sum(1), 2, 3, 4, 5, 6,)
  printNames("Jesus", "Rosa", "Val")
  fmt.Println(namedReturn(2))
}
