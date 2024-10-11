package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatValueFromFile(fileName string) (float64, error) {
	valueFromFile, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("Failed to find file")
	}

	valueTextString := string(valueFromFile)
	valueFloatValue, err := strconv.ParseFloat(valueTextString, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse the value from file")
	}

	return valueFloatValue, nil
}
func StoreFloatValueInfile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
