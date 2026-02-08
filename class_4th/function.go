package main
import ("fmt")

// func myMessage() {
//   fmt.Println("I just got executed!")
// }

// func main() {
//   myMessage() // call the function
//   myMessage() // call the function again

// }

// func deteteles(name string, age int) {
// 	fmt.Println("Name:", name)
// 	fmt.Println("Age:", age)
// }

// func main() {
// 	deteteles("Alice", 30)
// 	deteteles("Bob", 25)
// }


// func add(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	sum := add(5, 7)
// 	fmt.Println("Sum is:", sum)
// }

// recursive function

func factorial(x float64) ( y float64) {
	if x > 0 {
		y = x * factorial(x-1)
	} else {
		y = 1
	}
	return 
}

func main() {
	fmt.Println(factorial(4))
}





