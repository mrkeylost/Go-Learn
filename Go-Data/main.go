package main

import "fmt"

// pake type alias
// DRY principle
type floatMap map[string]float64

func (m floatMap) output () {
	fmt.Println(m)
}

func main() {
	// user := []string{}
	// params (tipe data, initial ukuran array, maximum kapasitas array)
	user := make([]string, 2, 5)

	user[0] = "Julia"
	user[1] = "Victor"

	user = append(user, "Max")
	user = append(user, "John")
	user = append(user, "Loki")
	user = append(user, "Nita")

	fmt.Println(user)

	courseRating := make(floatMap, 2)

	courseRating["Go"] = 4.7
	courseRating["React"] = 4.8
	courseRating["Vue"] = 4.8

	//fmt.Println(courseRating)
	courseRating.output()

	// looping an arrays
	for index, value := range user{
		fmt.Printf("%v is in index: %v\n", value, index)
	}

	// looping map
	for key, value := range courseRating{
		fmt.Printf("%v is in key: %v\n", value, key)
	}
}