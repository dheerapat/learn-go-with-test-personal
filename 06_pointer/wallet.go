package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFund = errors.New("error: insufficient balance")

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(money Bitcoin) {
	w.balance += money
}

func (w *Wallet) Withdraw(money Bitcoin) error {
	if money > w.balance {
		return ErrInsufficientFund
	}

	w.balance -= money
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
