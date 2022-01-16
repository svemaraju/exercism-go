package account

import "sync"

type Account struct {
	mu sync.Mutex
	balance int64
	active bool
}

// Balance returns the current balance of acc.
func (acc *Account) Balance() (int64, bool) {
	if (acc.active == false) {
		// account is currently closed.
		return 0, false
	}
	return acc.balance, true
}

// Close closes the account and returns payout.
func (acc *Account) Close() (int64, bool) {
	acc.mu.Lock()
	if (acc.active == false) {
		// account has already been closed.
		acc.mu.Unlock()
		return 0, false
	}
	acc.active = false
	acc.mu.Unlock()
	return acc.balance, true
}

// Deposit deposits/withdraws a valid amount, returns new balance.
func (acc *Account) Deposit(deposit int64) (int64, bool) {
	acc.mu.Lock()
	if (acc.active == false) {
		// account is currently closed.
		return 0, false
	}
	newBalance := acc.balance + deposit
	if (newBalance < 0) {
		// account balance will negative after deposit.
		acc.mu.Unlock()
		return acc.balance, false
	}
	acc.balance = newBalance
	acc.mu.Unlock()
	return acc.balance, true
}

// Open creates a new account with initial balance.
func Open(initialBalance int64) *Account {
	if (initialBalance < 0) {
		return nil
	}
	acc := Account{balance: initialBalance, active: true}
	return &acc
}


