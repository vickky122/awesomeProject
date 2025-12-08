package main

import "fmt"

func main() {
	revenue := inputText("Enter the revenue: ")
	expenses := inputText("Enter the expenses: ")
	taxrate := inputText("Enter the tax rate: ")

	ebt, profit, ratio := calculate(revenue, expenses, taxrate)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func inputText(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculate(revenue, expenses, taxrate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxrate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
