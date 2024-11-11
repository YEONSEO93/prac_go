package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood1 := []string{"gummy", "chocolate"}
	favFood2 := []string{"fruits", "beer"}

	yeonseo := person{"yeonseo", 20, favFood1}
	// yeonseo := person{"yeonseo", age:20, favFood1}
	fmt.Println(yeonseo)
	skye := person{name: "skye", age: 30, favFood: favFood2}
	// skye := person{name: "skye", 30, favFood: favFood2}
	fmt.Println(skye)
}
