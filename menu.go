package main

import "fmt"

func RunBank(account *Account) {
	exitMenu := false

	fmt.Println("Welcome to the Go bank!")

	for !exitMenu {
		fmt.Println("Please select an option:")

		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Balance")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
	}
}