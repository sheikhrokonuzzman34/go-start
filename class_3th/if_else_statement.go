// package main
// import ("fmt")

// func main() {
//   time := 20
//   if (time < 18) {
//     fmt.Println("Good day.")
//   } else {
//     fmt.Println("Good evening.")
//   }
// }



// func main() {
//   temperature := 14
//   if (temperature > 15) {
//     fmt.Println("It is warm out there")
//   } else {
//     fmt.Println("It is cold out there")
//   }
// }




// func main() {
//   time := 22
//   if time < 10 {
//     fmt.Println("Good morning.")
//   } else if time < 20 {
//     fmt.Println("Good day.")
//   } else {
//     fmt.Println("Good evening.")
//   }
// }

// func main() {
//   x := 30
//   if x >= 10 {
//     fmt.Println("x is larger than or equal to 10.")
//   } else if x > 20 {
//     fmt.Println("x is larger than 20.")
//   } else {
//     fmt.Println("x is less than 10.")
//   }
// }


package main
import ("fmt")

func main() {
   day := 5

   switch day {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }
}