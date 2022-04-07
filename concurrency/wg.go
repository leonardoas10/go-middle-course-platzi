package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // Suma una rutina
		go doSomething(i, &wg)
	}

	wg.Wait() // Espera que termine todo.
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done() // resta las rutinas agregadas en Add()
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}