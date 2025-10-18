package main

import "fmt"

func main() {
	age := 32 // regular

	// tanpa *, dia bakal ngambil address-nya
	// butuh dereferencing kalau mau ambil value
	// var agePointer *int
	agePointers := &age

	fmt.Println("Age:", *agePointers)

	// ganti value age dari 32 jadi 14 di address yang sama 
	editAgeToAdultYears(agePointers)
	fmt.Println("Adult Years:", age)
}

// dia gk bisa lakuin operasi pake pointers
func editAgeToAdultYears(age *int) int {
	// return *age - 18

	// data mutation
	// basically, dia ganti value data yang direference sama *age ini
	*age = *age - 18

	return *age
}