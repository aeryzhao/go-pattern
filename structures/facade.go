package main

import "fmt"

type account struct {
	name string
}

func newAccount(accountName string) *account {
	account := &account{
		name: accountName,
	}
	return account
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Account Name is incorrect")
	}
	println("Account verified")
	return nil
}

type wallet struct {
	balence int
}

func newWallet() *wallet {
	wallet := &wallet{
		balence: 0,
	}
	return wallet
}

func (w *wallet) creditBalance(amount int) {
	w.balence += amount
	println("Wallet balance added successfully")
}

func (w *wallet) debitBalance(amount int) error {
	if w.balence < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	println("Wallet balance is deducted")
	w.balence = w.balence - amount
	return nil
}

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	securityCode := &securityCode{
		code: code,
	}
	return securityCode
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("Security Code is incorrect")
	}
	println("SecurityCode verified")
	return nil
}

type notification struct{}

func (n *notification) sendWalletCreditNotification() {
	println("Sending wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	println("Sending wallet debit notification")
}

type ledger struct{}

func (s *ledger) makeEntry(accountID, txnType string, amount int) {
	println("Make ledger entry for accountId", accountID, "with txnType", txnType, "for amount", amount)
}

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

func newWalletFacade(accountID string, code int) *walletFacade {
	println("Starting create account")
	walletFacade := &walletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &notification{},
		ledger:       &ledger{},
	}
	println("Account created")
	return walletFacade
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func main() {
	walletFacade := newWalletFacade("abc", 1234)
	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		println("Error:", err.Error())
	}
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		println("Error:", err.Error())
	}
}
