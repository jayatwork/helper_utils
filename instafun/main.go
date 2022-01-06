package main

import (
	"gopkg.in/ahmdrz/goinsta.v2"
)

func main() {
	insta := goinsta.New("USERNAME", "PASSWORD")

	// Export your configuration
	// after exporting you can use Import function instead of New function.
	// insta, err := goinsta.Import("~/.goinsta")
	// it's useful when you want use goinsta repeatedly.
	insta.Export("~/.goinsta")
}
