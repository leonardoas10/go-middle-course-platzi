package main

import "fmt"

func sum(values ...int) int  {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (double int, triple int, quad int) { // return by value
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main()  {
	fmt.Println(sum(1,2))
	fmt.Println(getValues(5))
}