
## Go Array & Slice — Short Q&A

### Q1. What is a Go slice?

**Ans:**
A slice is a **dynamic collection** in Go.
Python equivalent: **list**

---

### Q2. What is a Go array?

**Ans:**
An array is a **fixed-size collection**.
➡️ Python has **no direct equivalent**.

---

### Q3. What does `len(slice)` mean?

**Ans:**
Number of elements currently in the slice.

```go
len([]int{1, 2, 3}) // 3
```

---

### Q4. What does `cap(slice)` mean?

**Ans:**
Maximum number of elements the slice can hold **before resizing**.

```go
cap([]int{1, 2, 3}) // 3
```

---

### Q5. Why check `cap` if slice size is dynamic?

**Ans:**
To **improve performance** and avoid repeated memory allocation.

---

### Q6. How to create a slice with predefined capacity?

**Ans:**

```go
s := make([]int, 0, 10) // len=0, cap=10
myslice := []int{1, 2, 3, 4,5,6}
fmt.Println(len(myslice), cap(myslice)) // 4 4

myslice = append(myslice, 9)
fmt.Println(len(myslice), cap(myslice)) // 5 8 (capacity doubled)
```
```txt
6 6
7 12
```

```go
func main() {
  temperature := 14
  if (temperature > 15) {
    fmt.Println("It is warm out there")
  } else {
    fmt.Println("It is cold out there")
  }
}

// right sentex

func main() {
  temperature := 14
  if (temperature > 15) {
    fmt.Println("It is warm out there")
  } 
  else {
    fmt.Println("It is cold out there")
  }
}
// worng sentex becose go line end bujle else oi statement buje na
```
