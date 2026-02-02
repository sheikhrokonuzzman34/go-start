package main

import "fmt"

func main() {
	var username string = "admin"
	fmt.Println("Username:", username)
	fmt.Printf("Type of username: %T\n", username)

	var age int = 30
	fmt.Println("Age:", age)
	fmt.Printf("Type of age: %T\n", age)

	var score uint16 = 255
	fmt.Println("Score:", score)
	fmt.Printf("Type of score: %T\n", score)

	var total int64 = 10000000000
	fmt.Println("Total:", total)
	fmt.Printf("Type of total: %T\n", total)

	var isActive bool = true
	fmt.Println("Is Active:", isActive)
	fmt.Printf("Type of isActive: %T\n", isActive)

	var height float64 = 5.900000000000
	fmt.Println("Height:", height)
	fmt.Printf("Type of height: %T\n", height)

	var pi float32 = 3.14
	fmt.Println("Pi:", pi)
	fmt.Printf("Type of pi: %T\n", pi)

	var initial rune = 'A'
	fmt.Println("Initial:", initial)
	fmt.Printf("Type of initial: %T\n", initial)

	var grade byte = 100
	fmt.Println("Grade:", grade)
	fmt.Printf("Type of grade: %T\n", grade)


	var temperature1 float32 = 36.6
	fmt.Println("Temperature:", temperature1)
	fmt.Printf("%T", temperature1)

	var temperature float32 = 36.6
	fmt.Println("Temperature:", temperature)
	fmt.Printf("%T\n", temperature)
}