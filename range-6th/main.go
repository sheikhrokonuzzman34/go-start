package main

import "fmt"

// range
// func main() {
// 	slice := []string{"Hello", "World", "Go", "Programming"}

// 	for i, value := range slice {
// 		fmt.Printf("Index: %d, Value: %s\n", i, value)
// 	}
// }

func main() {
	slice := []string{"Hello", "World", "Go", "Programming"}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, slice[i])
	}
}
