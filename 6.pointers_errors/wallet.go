package main

type Bitcoin int

type Wallet struct{
	balance Bitcoin
}

//the receiver type is *Wallet rather than Wallet which you can read as "a pointer to a wallet".
func (w *Wallet) Deposit(amount Bitcoin){
	w.balance += amount //auto destructured, can also write (*w).balance
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}