package main

import (
	"fmt"
)

type employee struct {
	name   string
	id     int
	age    int
	salary float32
}

func (e employee) getDetails() {
	fmt.Printf("Employee Name: %v\n", e.name)
	fmt.Printf("Employee ID: %v\n", e.id)
	fmt.Printf("Age of employee: %v\n", e.age)
	fmt.Printf("Check my bank bro: %v\n", e.salary)
}

func (e employee) getId() int {
	return e.id
}

func (e employee) getSalary() float32 {
	return e.salary
}

func main() {
	emp := employee{name: "jason", id: 1125, age: 42, salary: 500000.32}
	emp2 := employee{name: "ericka", id: 1124, age: 35, salary: 750000.58}
	emp.getDetails()
	fmt.Printf("I clearly make a lot of money - - %v\n", emp.getSalary())
	fmt.Printf("I'm new employee with ID - -  %v\n", emp.getId())

	//Next employee record
	emp2.getDetails()
	fmt.Printf("I clearly make a lot of money - - %v\n", emp2.getSalary())
	fmt.Printf("I'm new employee with ID - -  %v\n", emp2.getId())
}
