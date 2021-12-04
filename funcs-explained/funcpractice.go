package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 0000
}

func main() {

	var x []int = []int{1125,1100,1010,1382}
	fmt.Println(x)
	changeFirst(x)
	fmt.Println(x)

}