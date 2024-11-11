package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("let's test!")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLenght, up := lenAndUpper("yeonseo")
	fmt.Println(totalLenght, up)
}
