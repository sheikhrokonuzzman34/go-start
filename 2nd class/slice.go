package main

import ("fmt")

func main() { 
	var slice1 []int = []int{10, 20, 30, 40, 50}
	fmt.Println("Slice 1:", slice1)

	slice2 := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("Slice 2:", slice2)

	fmt.Println("Length of slice1:", len(slice1))
	fmt.Println("Capacity of slice1:", cap(slice1))
	slice1 = append(slice1, 60, 70)
	fmt.Println("Slice 1 after append:", slice1)


	myslice := []int{1, 2, 3, 4,5,6}
	fmt.Println(len(myslice), cap(myslice)) // 4 4

	myslice = append(myslice, 9)
	fmt.Println(len(myslice), cap(myslice)) // 5 8 (capacity doubled)



}

	