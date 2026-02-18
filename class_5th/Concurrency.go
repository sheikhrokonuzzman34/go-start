package main

import "fmt"
import "time"


func main() {
	go genarate("Hello")
	genarate("World")
}

func genarate(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}