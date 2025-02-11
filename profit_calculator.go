package profit_calculator

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	var EBT, EAT, ratio float64
	fmt.Println("Enter revenue in inr")
	fmt.Scan(&revenue)

	fmt.Println("Enter taxRater")
	fmt.Scan(&taxRate)

	fmt.Println("Enter expenses in inr")
	fmt.Scan(&expenses)

	EBT = revenue - expenses
	EAT = EBT * (1 - taxRate/100)
	ratio = EBT / EAT

	fmt.Printf("EBT: %.0f \nEAT: %.0f\nRatio: %v\n", EBT, EAT, ratio)

}

func printValue(val1 string) {
	fmt.Println(val1)
}
