package main

import "fmt"

type Account struct {
	balance float64 
}

func NewAccount(initialBalance float64) *Account {
	return &Account{initialBalance}
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
	fmt.Println("Deposited:", amount)
	fmt.Println("New balance:", a.balance)
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.balance {
		return fmt.Errorf("Insufficient funds")
	}
	a.balance -= amount
	fmt.Println("Withdrew:", amount)
	fmt.Println("New balance:", a.balance)
	return nil
}

func (a *Account) Balance() float64 {
	return a.balance
}

