package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	Hello()

	var price float64
	var rate float64
	var year float64

	//fmt.Printf("Enter the price: ")
	outputText("Enter the price:")
	fmt.Scan(&price)

	//fmt.Printf("Enter the rate: ")
	outputText("Enter the rate: ")
	fmt.Scan(&rate)

	//fmt.Printf("Enter the year: ")
	outputText("Enter the year: ")
	fmt.Scan(&year)

	famount, fRealValue := calculateFutureValues(price, rate, year)

	//amount := price * math.Pow(1+rate/100, year)
	//futureRealValue := amount / math.Pow(1+inflationRate/100, year)
	//fmt.Println(amount, futureRealValue)
	//fmt.Println(futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", famount)
	formattedRFV := fmt.Sprintf("Future value(adjusted for inflation: %.2f\n", fRealValue)
	//fmt.Printf("Future Value: %T\n%0.2f\nFuture value(adjusted for inflation): %0.2f\n", amount, amount, futureRealValue)
	//%T for amount, then two decimal places for amount, then two decimal places for futureRealValue

	fmt.Print(formattedFV, formattedRFV)
}
func outputText(text string) {
	fmt.Print(text)
}
func calculateFutureValues(price, rate, year float64) (famount float64, fRealValue float64) {
	famount = price * math.Pow(1+rate/100, year)
	fRealValue = famount / math.Pow(1+inflationRate/100, year)
	//return famount, fRealValue
	return
}
