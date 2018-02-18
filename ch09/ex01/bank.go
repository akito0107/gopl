package main

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)
var withdrawResults = make(chan bool)
var clear = make(chan struct{}) // for testing

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	withdraws <- amount
	return <-withdrawResults
}

func Clear() {
	clear <- struct{}{}
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			if balance < amount {
				withdrawResults <- false
				continue
			}
			balance -= amount
			withdrawResults <- true
		case <-clear:
			balance = 0
		}
	}
}

func init() {
	go teller()
}
