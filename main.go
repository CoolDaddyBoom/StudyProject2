package main

import (
	"errors"
	"fmt"
	"sync"
)

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Deposit(amount float64) {
	u.Balance += amount
}

func (u *User) Withdraw(amount float64) error {
	if u.Balance < amount {
		return errors.New("Insufficient funds")
	} else {
		u.Balance -= amount
		return nil
	}
}

type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}

type PaymentSystem struct {
	Users        map[string]*User
	Transactions []Transaction
}

func (ps *PaymentSystem) AddUser(u *User) {
	ps.Users[u.ID] = u
}

func (ps *PaymentSystem) AddTransaction(t Transaction) {
	ps.Transactions = append(ps.Transactions, t)
}

func (ps *PaymentSystem) ProcessingTransactions(t Transaction) error {

	if t.Amount < 1 {
		return errors.New("Incorrect value")
	}

	fromUser, ok := ps.Users[t.FromID]
	if !ok {
		return errors.New("Sender user not found")
	}
	toUser, ok := ps.Users[t.ToID]
	if !ok {
		return errors.New("Receiver user not found")
	}

	err := fromUser.Withdraw(t.Amount)
	if err != nil {
		return err
	}
	toUser.Deposit(t.Amount)
	return nil
}

func Worker(ps *PaymentSystem, ch <-chan Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 0
	for t := range ch {
		counter++
		err := ps.ProcessingTransactions(t)
		if err != nil {
			fmt.Printf("Transaction result: %v\n", err)
		} else {
			fmt.Println("Transaction result: is successful")
		}
	}
}

func main() {

	ps := &PaymentSystem{Users: make(map[string]*User), Transactions: []Transaction{}}
	user1 := &User{ID: "1", Name: "Alice", Balance: 100.0}
	user2 := &User{ID: "2", Name: "Bob", Balance: 100.0}

	ps.AddUser(user1)
	ps.AddUser(user2)

	t1 := Transaction{FromID: "1", ToID: "2", Amount: -1}
	t2 := Transaction{FromID: "3", ToID: "1", Amount: 1000}
	t3 := Transaction{FromID: "1", ToID: "2", Amount: 99}

	ps.AddTransaction(t1)
	ps.AddTransaction(t2)
	ps.AddTransaction(t3)

	var ch chan Transaction = make(chan Transaction, len(ps.Transactions))
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Worker(ps, ch, &wg)
	}

	for _, t := range ps.Transactions {
		ch <- t
	}

	close(ch)

	wg.Wait()

	fmt.Println("\nИтого:")
	fmt.Printf("У первого пользователя получилось %.2f\n", ps.Users["1"].Balance)
	fmt.Printf("У второго пользователя получилось %.2f\n", ps.Users["2"].Balance)
}
