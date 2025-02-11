package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "accountBalance.txt"

func main() {

	accountBalance, fileError := readFile()

	if fileError != nil {
		fmt.Println("Error occurred while reading balance")
		panic(fileError)
		return
	}

	fmt.Println("Welcome to bank")
	fmt.Println("1: to deposit")
	fmt.Println("2: to withdraw")
	fmt.Println("3: to check balance")
	fmt.Println("4: Exit")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Enter amount which is to be deposited")
		var depositedMoney float64
		fmt.Scan(&depositedMoney)
		accountBalance += depositedMoney
		fmt.Printf("Your account balance is %0.2f\n", accountBalance)
		writeIntoFile(accountBalance)
		return

	case 2:
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)
		accountBalance -= withdrawalAmount
		writeIntoFile(accountBalance)
		fmt.Printf("Your account balance is %0.2f\n", accountBalance)
		return

	case 3:
		fmt.Printf("Your account balance is %0.2f\n", accountBalance)
		return

	default:
		return

	}

}

func writeIntoFile(accountBalance float64) {
	balanceString := fmt.Sprint(accountBalance)
	os.WriteFile(fileName, []byte(balanceString), 0644)
	fmt.Println("Successfully stored account balance")
}

func readFile() (float64, error) {
	fileContentBytes, err := os.ReadFile(fileName)

	if err != nil {
		return 0, errors.New("Failed to find balance file")
	}

	fileContentString := string(fileContentBytes)
	balance, err := strconv.ParseFloat(fileContentString, 64)

	if err != nil {
		return 0, errors.New("failed to parse balance")
	}

	return balance, nil
}
