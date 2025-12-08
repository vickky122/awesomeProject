package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxrate float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxrate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxrate/100)
	ratio := ebt / profit

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
