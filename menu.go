package main

import "fmt"

func RunBank(account *Account) {
	exitMenu := false

	fmt.Println("Welcome to the Go bank!")

	for !exitMenu {
		fmt.Println("Please select an option:")

		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", account.Balance())
			readBalanceFromFile(account)
		case 2:
			fmt.Print("Enter amount to deposit: ")
			var amount float64
			fmt.Scan(&amount)
			account.Deposit(amount)
			writeBalanceToFile(account)
		case 3:
			fmt.Print("Enter amount to withdraw: ")
			var amount float64
			fmt.Scan(&amount)
			if err := account.Withdraw(amount); err != nil {
				fmt.Println(err)
			}
			writeBalanceToFile(account)
		case 4:
			fmt.Println("Thank you for using Go bank!")
			exitMenu = true
		default:
			fmt.Println("Invalid choice.")
		}

	}
}
