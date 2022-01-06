package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	htgotts "github.com/hegedustibor/htgo-tts"
)

func Initialize() {

	//For now todo input is being sourced from file
	var todo string = "todolist.txt"
	team := []string{"Tony", "Loren", "Raghu", "Baskar", "Jamie", "Antonio", "Tommy"}
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	reader, err := os.Open(todo)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		speech.Speak(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// This program is dependent upon MPLAYER being installed
func main() {
	//Greeting
	t := time.Now()
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak("Whatcha trying to accomplish today?" + t.String())

	Initialize()
	// Eventually package as mobile app to be deployed to ios or driod

	//MORE FEATURES

}
