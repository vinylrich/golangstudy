package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	balance int
	mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) {
	a.mutex.Lock()
	a.balance -= val
	a.mutex.Unlock()
}

func (a *Account) Deposit(val int) {
	a.mutex.Lock()
	a.balance += val //a.balance=a.balance-val
	a.mutex.Unlock()
	//RegA RegB
}
func (a *Account) Balance() int {
	return a.balance
}

var accounts []*Account

func Transfer(sender, receiver, money int) {

	time.Sleep(100 * time.Millisecond)

	accounts[sender].Widthdraw(money) //빼고 난 다음에 넣어서 오류 발생
	accounts[receiver].Deposit(money)
	fmt.Println("Transfer", sender, receiver, money)

}

func GetTotalBalance() int {

	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}

	return total
}

func RandomTransfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}
	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}
func PrintTotalBalance() {
	fmt.Printf("Total:%d\n", GetTotalBalance())
}
func main() {
	//같은 메모리를 두 개의 cpu가 접근하다보니 메모리가 꼬임
	//lock을 걸어서 한 자원을 다른 곳에서 접근을 못 하게 하는 것

	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
	}
	go func() {
		for {
			Transfer(0, 1, 100)
		}
	}()
	go func() {
		for {
			Transfer(1, 0, 100)
		}
	}()
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
