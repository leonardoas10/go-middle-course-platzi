package main

import (
	"fmt"
	"sync"
	"time"
)

//c := [][]
//c := [goRoutine1][]
//c := [goRoutine1][goRoutine2]
//c := [][goRoutine2]
//c := [goRoutine3][goRoutine2]

func main()  {
	c := make(chan int, 2)
	var wg sync.WaitGroup // []

	for i := 0; i < 10; i++ {
		c <- 3
		wg.Add(1) // Suma una rutina
		go doSomething(i, &wg, c) // La c, indica mas alla de pasarle el canal. 
		// Le esta dando un cupo con c <- 1 a la routine doSomething
	}

	wg.Wait() // Espera que termine todo.

	fmt.Println("leo")
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done() // resta las rutinas agregadas en Add()
	fmt.Printf("ID %d started\n", i)
	time.Sleep(4 * time.Second) // -----------Espera-------
	fmt.Printf("ID %d ended", i)
	<- c // Libera el cupon en el canal
}