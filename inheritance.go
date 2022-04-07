package main

import "fmt"

type Person struct {
	name string
	age int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "full time employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (ftEmployee TemporaryEmployee) getMessage() string {
	return "temporary employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessagr(p PrintInfo)  {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "name"
	ftEmployee.age = 5
	ftEmployee.id = 1
	fmt.Printf("%v", ftEmployee)
	tEmployee := TemporaryEmployee{}
	getMessagr(tEmployee)
	getMessagr(ftEmployee)
	// GetMessage(ftEmployee.Person)
}