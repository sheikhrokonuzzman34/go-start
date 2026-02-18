package main

import "fmt"

func conter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := conter()

	increment1 := conter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment1())
}
