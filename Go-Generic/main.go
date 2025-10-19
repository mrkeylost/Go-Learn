package main

import "fmt"

func main() {
	// misal disini, T bakal dijadiin int
	result := add(1, 2)

	fmt.Println(result)
}

// generic
// bisa pake buat type apapun
// basically, type nya bakal ditentuin pas function add ini dipanggil
// misal disini, typenya bisa int, float64, atau string
func add[T int | float64 | string](a, b T) T {
	return a + b
}

// func add(a, b interface{}) interface{} {

// 	aint, aIsInt := a.(int)
// 	bint, bIsInt := b.(int)

// 	if aIsInt && bIsInt {
// 		return aint + bint
// 	}

// 	afloat, aIsFloat := a.(float64)
// 	bfloat, bIsFloat := b.(float64)

// 	if aIsFloat && bIsFloat {
// 		return afloat + bfloat
// 	}

// 	astring, aIsString := a.(string)
// 	bstring, bIsString := b.(string)

// 	if aIsString && bIsString {
// 		return astring + bstring
// 	}

// 	return nil
// }
