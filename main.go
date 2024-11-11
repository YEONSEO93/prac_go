package main

import (
	"fmt"

	"github.com/YEONSEO93/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("yeonseo")
	fmt.Println(account) // This will use the account variable
}
