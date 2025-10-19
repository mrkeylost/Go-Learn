package maps

import "fmt"

type Product struct {
	id string
	title string
	price float64
}

func main() {
	websites := map[string]string{
		"Google":             "https://google.com",
		"Amazon Web Service": "https://aws.com",
	}

	fmt.Println(websites)

	// print spesific value from map
	fmt.Println(websites["Amazon Web Service"])

	// add map value
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	// delete value from map
	delete(websites, "Google")
	fmt.Println(websites)
}
