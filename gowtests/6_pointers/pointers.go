package pointers

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) { // In golang method arguments are copied hence pointers are used to update by pointing at the value
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
