package main
// Ejecutar go test para ejecutar este archivo
// Ejecutar go test --cover para ver la cobertura del codigo en la terminal
// Ejecutar go test --coverprofile coverage.out para lanzar las metricas de la cobertura a un archivo
// Ejecutar go tool cover --func=coverage.out para ver la cobertura en la terminal
// EJecutar go tool cover --html=coverage.out para ver las metricas en el navegador
// Ejecutar go test --cpuprofile=cpu.out para ver las metricas de cpu (profiling) en un archivo
// Ejecutar go tool pprof cpu.out para er el resumen de sos del cpu (dentro ejecutar top, web, pdf)

import "testing"

func TestSumSimple(t *testing.T) {

	expectedResult := 10
	result := Sum(5, 5)
	if expectedResult !=  result {
		t.Errorf("Error en la suma, es obtuvo %d, se esperaba %d", result, expectedResult)
	}

}

func TestSumTableCase(t *testing.T) {
	tables := []struct {
		a int
		b int
		result int
	} {
		{
			a: 5,
			b: 5,
			result: 10,
		},
		{
			a: 1,
			b: 2,
			result: 3,
		},
		{
			a: 20,
			b: 25,
			result: 45,
		},
	}

	for _, value := range tables {
		result := Sum(value.a, value.b)
		if result != value.result {
			t.Errorf("Error en la suma, se obtuvo %d, se esperaba %d", result, value.result)
		}
	}
}


func TestMaxNum(t *testing.T) {
	tables := []struct{
		a int
		b int
		max int
	} {
		{3,2,3},
		{5,4,5},
		{10,14,14},
	}

	for _, value := range tables {
		max := MaxNum(value.a, value.b)

		if max != value.max {
			t.Errorf("Error en MaxSum, se obtuvo %d, se esperaba %d", max, value.max)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tables := [] struct {
		a int
		fib int
	} {
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, value := range tables {
		fib := Fibonacci(value.a)
		if fib != value.fib {
			t.Errorf("Error en fibonacci, se obtuvo %d, se esperaba %d", fib, value.fib)
		}
	}
}