package main

import "fmt"

func main() {
	names := []string{"yeonseo", "ko", "ys"}
	names = append(names, "slice")
	fmt.Println(names)
}
