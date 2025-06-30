package main

import "fmt"

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Deposit(amount float64) {
	u.Balance += amount
}

func (u *User) Withdraw(amount float64) bool {
	if u.Balance-amount < 0 {
		return false
	} else {
		u.Balance -= amount
		return true
	}
}

func main() {

	user1 := &User{ID: "1", Name: "Alice", Balance: 100.0}
	user2 := &User{ID: "2", Name: "Bob", Balance: 200.0}

	user1.Deposit(100.1)
	fmt.Printf("%s's balance after a deposit: %.2f\n", user1.Name, user1.Balance)
	user2.Withdraw(199.9)
	fmt.Printf("%s's balance after a withdraw: %.2f\n", user2.Name, user2.Balance)
	user1.Withdraw(199.9)
	fmt.Printf("%s's balance after a withdraw: %.2f\n", user1.Name, user1.Balance)
	user2.Deposit(100)
	fmt.Printf("%s's balance after a deposit: %.2f\n", user2.Name, user2.Balance)

	fmt.Println(user1, user2)
}
