package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("******************************************************")
	fmt.Println("I'm needing you YEAH YOU to input a number from 0 - 10")
	fmt.Println("******************************************************\n\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fromConsole := scanner.Text()
		fmt.Println(fromConsole)
		toInt, _ := strconv.Atoi(fromConsole)

		if _, err := strconv.Atoi(fromConsole); err == nil {
			fmt.Printf("%q looks like a number.\n", fromConsole)
			if toInt >= 1 && toInt < 10 {
				fmt.Printf("You've done well inputting a real number %v", toInt)
				if err := scanner.Err(); err != nil {
					fmt.Fprintln(os.Stderr, "reading standard input:", err)
				}
				break
			}
		}

		fmt.Printf("Reminder we need an int 0 thru 10 \n You entered: %v", fromConsole)

	}

}
