package main

import (
	"bufio"
	"example/notes/notes"
	"example/notes/todo"
	"fmt"
	"os"
	"strings"
)

// interface
// berguna untuk mendefine method yang ada disini + value yang direturn
type Saver interface {
	Save() error
}

// type Displayer interface {
// 	Display()
// }
// embedded interface
type Outputtable interface {
	Saver
	Display()
}

// special interfaces
// type General any

func main() {
	printSomething("Welcome To Notes App!")

	title, content:= getNoteData()
	todoText := gettingUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	noteData, err := notes.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	// todo.DisplayTodo()
	outputData(todo)

	// if err != nil {
	// 	return
	// }

	// noteData.DisplayNote()
	outputData(noteData)

	// if err != nil {
	// 	return
	// }
}

// interface implementation
func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Failed when saving")
		return err
	}

	fmt.Print("Save Data Succeeded")
	return nil
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}

// semua value bisa dipakai disini
func printSomething(value interface{}) {	
	// type switches
	// conditional statement berdasarkan type dari value
	// switch value.(type) {
	// case int: 
	// 	fmt.Println("integer :", value)
	// case float64:
	// 	fmt.Println("Float64 :", value)
	// case string:
	// 	fmt.Println(value)
	// }

	// kalau type data diluar ini, dia bakal ignore
	
	// fmt.Println(value)
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
