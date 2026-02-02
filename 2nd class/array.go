package main


import ("fmt")

func main() { 
	var arr1 [3]int = [3]int{10, 20, 30}
	fmt.Println("Array 1:", arr1)
	fmt.Printf("Type of arr1: %T\n", arr1)

	var arr2 = [4]string{"Go", "Python", "Java", "C++"}
	fmt.Println("Array 2:", arr2)
	fmt.Printf("Type of arr2: %T\n", arr2)

	var arr3 = [...]float64{3.14, 1.618, 2.718}
	fmt.Println("Array 3:", arr3)

	var arr4 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Array 4:", arr4)

	arr6 := [5]int{4,5,6,7,8}
	
	arr5 := [5]int{100, 200, 300, 400, 500}
	fmt.Println("Array 5:", arr5)
}