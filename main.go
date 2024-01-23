package main

import "fmt"

func main() {
	account := NewAccount(0) // Start with a zero balance

	// Load balance from file
	if err := readBalanceFromFile(account); err != nil {
		account.Deposit(1000) // Set a default balance if the file doesn't exist or can't be read
	}

	RunBank(account)

	// Save balance to file when exiting
	if err := writeBalanceToFile(account); err != nil {
		fmt.Println("Error saving balance:", err)
	}
}
