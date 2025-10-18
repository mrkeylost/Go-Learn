package main

import (
	"bufio"
	"example/notes/notes"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content:= getNoteData()

	noteData, err := notes.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	noteData.DisplayNote()
	err = noteData.SaveNote()

	if err != nil {
		fmt.Println("Failed when saving")
		return
	}

	fmt.Print("Save Note Succeeded")
}

func getNoteData() (string, string) {
	title := gettingUserInput("Note Title: ")
	content := gettingUserInput("Note Content: ")

	return title, content
}

func gettingUserInput(promptText string) string {
	fmt.Print(promptText)
	// var value string

	// scanln agak kurang untuk case input multiple words
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
