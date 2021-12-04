package main

import (
	"fmt"
)

func insertIntoMemLocation(newvalue *string) string {
	var temp string
	temp = *newvalue
	temp = "NOToutchere"
	return temp
}

func main() {
	value := "outchere"
	fmt.Println("this is simply the value", value, "\n and this is the memory location", &value)
	fmt.Println("\n Now lets see the switch on value based on func return", insertIntoMemLocation(value))



}