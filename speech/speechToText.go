package main

import (
	htgotts "github.com/hegedustibor/htgo-tts"
)

func main() {
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak("You are an awesome golang programmer.")
}
