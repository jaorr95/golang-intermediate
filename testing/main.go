package main


func Sum(a, b int)  int {
	return a + b
}

func MaxNum(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func Fibonacci(a int) int {
	if a <= 1 {
		return a
	}

	return Fibonacci(a - 1) + Fibonacci( a - 2)
}