package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// harus dimulai pake capital kalo mau diexport
func GetFloatFromFile(fileName string) (float64, error) {
	// _ dipake buat mau handle value, tapi belom mau dipake
	// kalo pake variable, bakal ada error karena unused
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to read file.")
	}

	valueText := string(data)
	// 2 params
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func SaveFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)

	// 0644 : file permission
	os.WriteFile(fileName, []byte(valueText), 0644)
}