package main

import (
	"errors"
	"fmt"
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

func main() {

	ps := &PaymentSystem{Users: make(map[string]*User), Transactions: []Transaction{}}
	user1 := &User{ID: "1", Name: "Alice", Balance: 100.0}
	user2 := &User{ID: "2", Name: "Bob", Balance: 100.0}

	ps.AddUser(user1)
	ps.AddUser(user2)

	t1 := Transaction{FromID: "1", ToID: "2", Amount: -1}
	t2 := Transaction{FromID: "3", ToID: "1", Amount: 1000}
	t4 := Transaction{FromID: "1", ToID: "2", Amount: 99}

	ps.AddTransaction(t1)
	ps.AddTransaction(t2)
	ps.AddTransaction(t4)

	for i, t := range ps.Transactions {
		err := ps.ProcessingTransactions(t)
		if err != nil {
			fmt.Printf("Transaction number %d: %v\n", i+1, err)
		} else {
			fmt.Printf("Transaction number %d: Is successful\n", i+1)
		}
	}

	fmt.Println("\nИтого:")
	fmt.Printf("У первого пользователя получилось %.2f\n", ps.Users["1"].Balance)
	fmt.Printf("У второго пользователя получилось %.2f\n", ps.Users["2"].Balance)
}
