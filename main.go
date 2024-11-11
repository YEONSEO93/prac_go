package main

import (
	"fmt"

	"github.com/YEONSEO93/learngo/accounts"

	"log"
)

func main() {
	account := accounts.NewAccount("yeonseo")
	fmt.Println(account)
	account.Deposit(10)
	// fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
}
