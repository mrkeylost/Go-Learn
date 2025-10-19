package lists

import "fmt"

func main() {
	// dynamic list pake slices
	prices := []float64{10.99, 8.99}

	// bikin slices baru
	prices = append(prices, 11.99, 5.99, 2.39, 9.59)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 45.99, 20.59}
	// spread operator
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {

// 	// declare a variabel with array data types
// 	var productNames [4]string = [4]string{"A Book"}

// 	productNames[2] = "A Carpet"
// 	// array
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	fmt.Println(productNames[0])

// 	// slice 
// 	// mirip python => ngambil elemen spesifik di array
// 	// [element pertama : element terakhir - 1]
// 	featuredPrices := prices[1:]
// 	// kalo angka pertama kosong, berarti by default mulai dari index 0
// 	// featuredPrices := prices[:3]
// 	// kalo angka kedua kosong, berarti berakhir sampe index terkahir
// 	// featuredPrices := prices[1:]

// 	// ini gak bisa, dibawah index terkecil atau diatas index terakhir
// 	// prices[-1:3] || prices[1:7]

// 	highlightPrices := featuredPrices[:1]

// 	fmt.Println(featuredPrices, highlightPrices)
// 	// len : banyak elemen dalam sebuah array
// 	// cap : capacity array

// 	// !! fun fact
// 	// ketika akses kapasitas array, dia gk bisa baca element kekiri, tapi gk kekanan 
// 	fmt.Println(len(highlightPrices), cap(highlightPrices))

// 	// kayak gini, dia bisa select ke arah kanan
// 	highlightPrices = highlightPrices[:3]
// 	fmt.Println(len(highlightPrices), cap(highlightPrices))
// }