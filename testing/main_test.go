package main

import "testing"

// go tool coverprofile=<name-file> **genera un archivo con los cover
//go tool cover -html= <name-file> proyecta en el browser el cover

func TestSum(t *testing.T)  {
	// total := Sum(5,5)

	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	// }
	tables := []struct {
		a int
		b int
		n int
	}{
		{1,2,3},{2,2,4},{5,4,9},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4,2,4},{2,3,3},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("Get Max was incorrect, got %d, expected %d", max, item.n)
		}
	}
}

func TestFibonacci(t *testing.T)  {
	tables := []struct {
		a int
		n int
	} {
		{1,1},
		{8,21},
		{50, 12586269025},
	}
	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d, expected %d", fib, item.n)
		}
	}
}