package tddgo

import "fmt"

type (
	Bitcoin int
	Wallet  struct {
		ballance Bitcoin
	}
)

var ErrInsufficientFunds = fmt.Errorf("cannot withdraw, insufficient funds")

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(quantidade Bitcoin) {
	w.ballance += quantidade
}

func (w *Wallet) Ballance() Bitcoin {
	return w.ballance
}

func (w *Wallet) Withdraw(quantidade Bitcoin) error {
	if quantidade > w.ballance {
		return ErrInsufficientFunds
	}
	w.ballance -= quantidade
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
