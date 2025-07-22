# Variable Types in Go

In programming, variable types define how values are stored, processed, and manipulated within code. Golang (Go) offers a rich set of data types, allowing developers to write efficient and maintainable code. This article explores basic, composite, and reference data types in Go to help you understand their functionality and usage.

---

## What Are Variable Types?

In Go, a variable type is a classification that determines the kind of value a variable can hold. Understanding variable types enables you to:

- ✅ Optimize code performance  
- ✅ Ensure data correctness  
- ✅ Avoid common programming errors  

Variable types form the foundation for storing, modifying, and retrieving data in a structured and efficient manner.

---

## Basic Data Types in Go

Go provides several fundamental data types that cater to everyday programming needs.

### 1. Integer Types

Go supports signed and unsigned integers with different bit sizes to accommodate various ranges of values.

| Type   | Bits | Range                                    |
|--------|------|------------------------------------------|
| int8   | 8    | -128 to 127                              |
| int16  | 16   | -32,768 to 32,767                        |
| int32  | 32   | -2,147,483,648 to 2,147,483,647          |
| int64  | 64   | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| uint8  | 8    | 0 to 255                                 |
| uint16 | 16   | 0 to 65,535                              |
| uint32 | 32   | 0 to 4,294,967,295                       |
| uint64 | 64   | 0 to 18,446,744,073,709,551,615          |

**Example:**

```go
var age int32 = 25
var score uint16 = 100
```

---

### 2. Floating-Point Types

Floating-point numbers are used for numbers with decimals.

| Type    | Precision                  |
|---------|----------------------------|
| float32 | Up to 6–8 decimal places   |
| float64 | Up to 15–16 decimal places |

**Example:**

```go
var pi float64 = 3.141592653589793
```

---

### 3. Boolean Type

The Boolean type represents a value of `true` or `false`. It’s commonly used in conditional logic and loops.

**Example:**

```go
var isReady bool = true
```

---

### 4. String Type

A string is a sequence of characters used to store textual information. Strings in Go are immutable.

**Example:**

```go
var message string = "Hello, Go!"
```

---

## Composite Data Types in Go

Composite data types allow combining multiple values, either of the same or different types, into a single structure.

### 1. Arrays

Arrays are fixed-length sequences of elements of the same type.

**Example:**

```go
var arr [5]int // Array with 5 integers
arr[0] = 10
```

---

### 2. Slices

A slice is a dynamic array that can grow or shrink in size.

**Example:**

```go
var s []int = []int{1, 2, 3, 4, 5}
```

---

### 3. Structs

A struct is a collection of fields that can have different data types.

**Example:**

```go
type Person struct {
    Name string
    Age  int
}
var p = Person{Name: "Alice", Age: 25}
```

---

### 4. Maps

A map is a collection of key-value pairs.

**Example:**

```go
var m map[string]int = make(map[string]int)
m["age"] = 30
fmt.Println(m["age"]) // Output: 30
```

---

### 5. Channels

Channels enable communication between Goroutines.

**Example:**

```go
ch := make(chan int)
go func() {
    ch <- 42
}()
value := <-ch
fmt.Println(value) // Output: 42
```

---

## Reference Types in Go

Reference types store memory addresses instead of values directly.

### 1. Pointers

Pointers hold the memory address of a value.

**Example:**

```go
var x int = 10
var ptr *int = &x
fmt.Println(*ptr) // Output: 10
```

---

## Why Understanding Variable Types Is Important

By mastering the differences between basic, composite, and reference types, you can:

- Write optimized, clear, and concise code  
- Avoid runtime errors and bugs  
- Take full advantage of Go’s features  

---

## Conclusion

Mastering variable types in Go is essential for writing efficient, reliable, and maintainable code. By understanding and correctly using basic, composite, and reference types, you can:

- ✅ Optimize your programs  
- ✅ Avoid common errors  
- ✅ Make the most of Go’s powerful features

