package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "profit.txt"

func main() {

	var EBT, EAT, ratio float64

	revenue, err := readValues("Enter revenue in inr")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := readValues("Enter taxRate")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := readValues("Enter expenses in inr")

	if err != nil {
		fmt.Println(err)
		return
	}

	EBT = revenue - expenses
	EAT = EBT * (1 - taxRate/100)
	ratio = EBT / EAT

	fmt.Printf("EBT: %.0f \nEAT: %.0f\nRatio: %v\n", EBT, EAT, ratio)

	storeData(EBT, EAT, ratio)

}

func storeData(EBT float64, EAT float64, ratio float64) {
	eBTString := strconv.FormatFloat(EBT, 'f', 2, 64)
	eATString := strconv.FormatFloat(EAT, 'f', 2, 64)
	ratioString := strconv.FormatFloat(ratio, 'f', 2, 64)
	fileContent := "	EBT: " + eBTString + "	EAT: " + eATString + "	ratio: " + ratioString
	err := os.WriteFile(fileName, []byte(fileContent), 0644)
	if err != nil {
		fmt.Println("Failed to write data in file")
		fmt.Println(err)
		return
	}
	fmt.Println("Date written in file successfully")
}

func printValue(val1 string) {
	fmt.Println(val1)
}

func readValues(msg string) (float64, error) {
	fmt.Println(msg)
	var scannedValue float64
	_, err := fmt.Scanf("%f", &scannedValue)
	if err != nil {
		fmt.Println("Invalid value")
		return scannedValue, errors.New("Failed to read value")
	}
	return scannedValue, nil
}
