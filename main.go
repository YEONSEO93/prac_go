package main

import "fmt"

func main() {
	ys := map[string]string{"name": "ys", "age": "123"}
	for key, value := range ys {
		fmt.Println(key, value)
	}
}
