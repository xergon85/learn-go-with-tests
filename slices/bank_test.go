package slices

import (
	"testing"

	"github.com/xergon85/learn-go-with-tests/generics"
)

func TestBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}
	generics.AssertEqual(t, BalanceFor(transactions, "Riya"), 100)
	generics.AssertEqual(t, BalanceFor(transactions, "Chris"), -75)
	generics.AssertEqual(t, BalanceFor(transactions, "Adil"), -25)
}

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}
	)

	transactions := []Transaction{
		NewTransaction(chris, riya, 100),
		NewTransaction(adil, chris, 25),
	}

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	generics.AssertEqual(t, newBalanceFor(riya), 200)
	generics.AssertEqual(t, newBalanceFor(chris), 0)
	generics.AssertEqual(t, newBalanceFor(adil), 175)
}
