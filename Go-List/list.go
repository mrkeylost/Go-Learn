package main

import "fmt"

type Product struct {
	title string
	id string
	price float64
}

func main() {
	hobby := [3]string{"Gaming", "Streaming", "Walking"}

	// #1
	fmt.Println(hobby)

	// #2
	fmt.Println(hobby[0])
	fmt.Println(hobby[1:])

	// #3
	priorityHobby := hobby[:2]
	fmt.Println(priorityHobby)

	// #4
	// Gak bisa pake [1:] karena dia ambil dari priority skrng
	priorityHobby = priorityHobby[1:3]
	fmt.Println(priorityHobby)

	// #5
	goals := []string{"Get good", "master API"}
	fmt.Println(goals)

	// #6
	goals[1] = "Fundamentals Comprehension"
	goals = append(goals, "Think like a programmer") 
	fmt.Println(goals)

	// #7
	// myProduct := Product {
	// 	title: "Catalog",
	// 	id: "001",
	// 	price: 345.0,
	// 	list: []string{"Book", "Pencil"},
	// }
	// myProduct.list = append(myProduct.list, "Eraser")
	// fmt.Println(myProduct)

	products := []Product{
		{
			title: "Book", 
			id: "001", 
			price: 10.0,
		}, 
		{
			title: "Pencil",
			id: "002",
			price: 7.0,
		},
	}
	fmt.Println(products)

	newProd := append(products, Product{
		title: "Eraser", id: "003", price: 5.0,
	})
	fmt.Println(newProd)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.