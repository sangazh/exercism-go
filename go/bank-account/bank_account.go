package account

import "sync"

type Account struct {
	balance int64
	status  bool
	mux     sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit, status: true}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if !a.status {
		return 0, false
	}

	a.status = false
	return a.balance, true
}
func (a *Account) Balance() (balance int64, ok bool) {
	if !a.status {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if !a.status {
		return 0, false
	}

	if amount < 0 {
		if a.balance+amount < 0 {
			return a.balance, false
		}
	}
	a.balance += amount
	return a.balance, true
}
