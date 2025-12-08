package main

import (
	"awesomeProject/bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		//panic("Can't continue. Sorry")
	}
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
