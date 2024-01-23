package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(account *Account) error {
	file, err := os.Create("balance.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%f", account.balance)
	return err
}

func readBalanceFromFile(account *Account) error {
	file, err := os.Open("balance.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fscanf(file, "%f", &account.balance)
	return err
}
