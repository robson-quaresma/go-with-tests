package pointerserrors

import (
	"errors"
	"fmt"
)

// New: create a self error
// with var keyword we can create global values in the package
var ErrInsufficientFounds = errors.New("cannot withdraw, insufficient funds")

// New: Go allows us to create new types from existing ones.
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFounds
	}
	w.balance -= amount
	return nil
}
