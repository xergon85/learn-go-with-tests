package slices

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	calculateBalance := func(currentBalance float64, transaction Transaction) float64 {
		if transaction.From == name {
			return currentBalance - transaction.Sum
		}
		if transaction.To == name {
			return currentBalance + transaction.Sum
		}

		return currentBalance
	}

	return Reduce(transactions, calculateBalance, 0.0)
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{from.Name, to.Name, sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}

	return a
}
