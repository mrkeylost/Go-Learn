package main

import (
	"fmt"
)

// type bisa dipake buat semua data type, bukan cuma struct

type CustomString string

// bisa diapke buat built in method
func (text CustomString) log() {
	fmt.Println(text)
}

func main() {
	var name CustomString

	name.log()
}