package main

import "fmt"

type Account struct {
	name string
}

type Wallet struct {
	balance float64
}

type Notification struct{}

func (a *Account) Validade(name string) bool {
	if a.name == name {
		fmt.Println("Account validated")
		return true
	}
	fmt.Println("Account validation failed: ", name)
	return false
}

func (w *Wallet) CheckBalance(amount float64) bool {
	if w.balance >= amount {
		fmt.Println("Sufficient balance available")
		return true
	}
	fmt.Println("Insufficient balance")
	return false
}

func (w *Wallet) Deduct(amount float64) {
	w.balance -= amount
	fmt.Printf("Amount %.2f deducted. New balance  %.2f\n", amount, w.balance)
}

func (n *Notification) Send(message string) {
	fmt.Println("Notification sent: ", message)
}

type BankFacade struct {
	account      *Account
	wallet       *Wallet
	notification *Notification
}

func NewBankFacade(name string, balance float64) *BankFacade {
	return &BankFacade{
		account:      &Account{name},
		wallet:       &Wallet{balance},
		notification: &Notification{},
	}
}

func (b *BankFacade) MakeTransaction(accountName string, amount float64) {
	fmt.Println("\nStarting Transaction...")

	if !b.account.Validade(accountName) {
		return
	}

	if !b.wallet.CheckBalance(amount) {
		return
	}

	b.wallet.Deduct(amount)
	b.notification.Send(fmt.Sprintf("Transaction of %.2f successful", amount))
	fmt.Println("Transaction completed successfully.")
}

func main() {
	bank := NewBankFacade("Eduardo", 500)

	bank.MakeTransaction("Eduardo", 200)
	bank.MakeTransaction("Test", 100)

}
