package main

import (
	"github.com/YEONSEO93/learngo/mydict"

	"fmt"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First Word"}
	definition, err := dictionary.Search("second")
	// definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
