package pointers

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (w *Wallet) Deposit(amount Bitcoin) { // In golang method arguments are copied hence pointers are used to update by pointing at the value
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error { // In golang method arguments are copied hence pointers are used to update by pointing at the value
	if amount > w.balance {
		return errors.New("oh no")
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
