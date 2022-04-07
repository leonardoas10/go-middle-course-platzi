package main

import "fmt"

type Employee struct {
	id int
	name string
}

func (e *Employee) SetId(id int)  {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func main() {
	e := Employee{}
	fmt.Printf("%v", e)
	e.id = 0
	e.name = "leo"
	e.SetId(5)
	e.SetName("leoss")
	fmt.Printf("%v", e.GetId())
}