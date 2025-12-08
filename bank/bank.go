package main

import (
	"fmt"
	"os"
	"strconv"
)

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile("balance.txt")
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
func main() {
	var accountBalance float64 = getBalanceFromFile()

	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			var depositBalance float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&depositBalance)
			if depositBalance < 0 {
				fmt.Println("Invalid deposit amount!")
				//return
				continue
			}
			accountBalance += depositBalance
			fmt.Println("Your balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			var withdrawbalance float64
			fmt.Print("Enter the money to withdraw: ")
			fmt.Scan(&withdrawbalance)
			if withdrawbalance < 0 {
				fmt.Println("Invalid withdraw amount!")
				//return
				continue
			}
			if withdrawbalance > accountBalance {
				fmt.Println("Invalid withdraw amount as you cannot withdraw more money than your account balance")
				//return
				continue
			}

			accountBalance -= withdrawbalance
			fmt.Println("Your balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 4:

			fmt.Println("Thank you!")
			return
		default:
			fmt.Println("Invalid choice!")
			fmt.Println("Thank you for using Go Bank!")

			//continue
			return
		}
	}

}
