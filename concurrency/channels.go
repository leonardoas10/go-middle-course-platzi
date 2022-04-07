package main

import "fmt"

func main()  {
	c := make(chan int, 2)
	c <- 4
	c <- 3
	fmt.Println(len(c), cap(c))
	// fmt.Println(<-c)
	close(c)
	for v := range c {
		fmt.Println(v)
	}

	// select

	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <- email1:
			fmt.Println("Email recibido 1", m1)
		case m2 := <- email2:
		fmt.Println("Email recibido 2", m2)
		}
	
	}
}

func message(text string, c chan<- string)  {
	c <- text
}