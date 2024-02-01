package account

import "sync"

// Define the Account type here.
type Account struct {
	balance int64
	isOpen bool
	mu *sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{ balance: amount, isOpen: true, mu: &sync.Mutex{}}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.isOpen {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.isOpen || a.balance + amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isOpen {
		bal := a.balance
		a.balance = 0
		a.isOpen = false
		return bal, true
	}
	return 0, false
}
