
# 📦 Go Packages & Code Organization

## 🔑 What is a Package?
A **Go package** is a group of `.go` files in the same folder that are compiled together. It helps organize code for **reusability** and **maintainability**.

### ✅ Key Features
- **Declaration**: All files in a folder must use the **same package name**.
- **Reusability**: Avoids code duplication.

---

## 🧰 Types of Packages

### 1. **Executable Packages**
- Used to build **standalone programs**
- Must be named `main`
- Must contain a `main()` function

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

**Output:**
```
Hello, World!
```

---

### 2. **Library Packages**
- Contain **reusable code** (functions, variables)
- **Do not include** `main()` function
- Are imported into other packages

```go
// File: math/math.go
package main

import (
    "fmt"
    "math"
)

func main() {
    result := math.Sqrt(25)
    fmt.Println("The square root is:", result)
}
```

**Output:**
```
The square root is: 5
```

---

## 📁 Package Structure & Naming

### 🔹 1. Directory Structure

Project structure matters. Packages are identified by folder paths.

```
yourproject/
├── main.go
└── internal/
    └── wow/
        └── wow.go
```

### 🔹 2. Importing Packages

Use import paths **relative to root**:

```go
import "yourproject/internal/wow"
```

### 🔹 3. Using Functions

Use the **package name as a prefix**:

```go
wow.SomeFunction()
```

---

## ⚙️ Special Functions

### `init()` Function
- Automatically runs **when a package is imported**
- Great for **initial setup**

```go
package wow

import "fmt"

func init() {
    fmt.Println("Wow package initialized!")
}

func SomeFunction() {
    fmt.Println("Some function in wow package")
}
```

**Output when imported:**
```
Wow package initialized!
```

---

## 🔐 Variable Scoping

| Name Format   | Scope                        |
|---------------|------------------------------|
| lowercase     | **Private** (within package) |
| Uppercase     | **Public** (outside package) |

```go
package math

var x = 10 // private
var X = 20 // public
```

---

## ✅ Best Practices

- ✅ Use **meaningful names**
- ✅ Break into **small, modular packages**
- ✅ Group similar functions together
- ✅ Use `internal/` folder for **private packages**
- ✅ **Comment your code** for clarity

---

## 🧾 Conclusion

Go packages are powerful tools to **organize**, **reuse**, and **encapsulate** code.  
Understanding executable vs library packages, `init()` behavior, and scoping rules makes your code **clean**, **scalable**, and **maintainable**.
