
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


Learn

pprof

CPU vs memory profiling

Allocation reduction

Escape analysis basics

Do

Profile your API

Optimize a slow endpoint

Outcome

You understand where performance actually comes from

Week 10 – Runtime & Advanced Topics

Learn

Go scheduler (G-M-P model)

Garbage collector behavior

sync/atomic

Struct memory layout

Do

Measure GC impact

Fix false sharing issues

Outcome

You reason about Go under load

Weeks 11–12: Production & Mastery
Week 11 – Testing, Reliability & Security

Learn

Race detector (-race)

Fuzz testing

Secure coding practices

TLS basics

Do

Add fuzz tests

Run chaos tests

Fix discovered issues

Outcome

Your code survives hostile conditions

Week 12 – Capstone Project

Build

A production-grade Go service:

REST or gRPC API

Concurrency

Tests

Metrics & logging

Graceful shutdown

Deploy (Optional but Powerful)

Dockerize it

Add basic CI

Deploy to a cloud VM or Kubernetes

Outcome

You have a portfolio-ready Go project
