
# 📦 Imports & Exports in Go

Managing external packages and dependencies is a critical part of Go development. The **module system** helps organize and lock dependencies using the `go.mod` file.

---

## 📄 What is `go.mod`?
The `go.mod` file is the **version manager** for Go projects. It keeps track of external packages and ensures consistency.

### 🔑 Key Features
- Tracks versions of external packages
- Ensures reproducible builds
- Works with commands like `go build`, `go test`
- Without it, Go commands may fail or behave unpredictably

---

## 🛠️ How to Initialize `go.mod`

### 1. Navigate to Your Project Directory

```bash
cd your-project-folder
```

### 2. Initialize `go.mod`

```bash
go mod init <module-name>
# Example:
go mod init models
```

This creates a `go.mod` file specifying the module name and Go version.

---

## 📥 Importing External Packages

Go makes it simple to import and manage external packages. The `go.mod` file automatically updates dependencies.

### 🔄 Example: Google UUID Package

#### 1. Install the Package

```bash
go get github.com/google/uuid
```

> This updates your `go.mod` file and downloads the required package.

#### 2. Use the Package in Code

```go
package main

import (
    "fmt"
    "github.com/google/uuid"
)

func main() {
    newUUID := uuid.New()
    fmt.Println(newUUID)
}
```

#### 3. Run the Program

```bash
go run main.go
```

---

## 🧰 Useful Go Module Commands

### ✅ 1. `go mod init`

Initializes a new `go.mod` file.

```bash
go mod init <module-name>
```

### 🧹 2. `go mod tidy`

- Removes unused dependencies
- Adds any missing dependencies

```bash
go mod tidy
```

> Keeps your module clean and conflict-free.

### 📋 3. `go list`

Lists all packages/modules currently in use.

```bash
go list
```

---

## 🧾 Conclusion

Go's module system—especially the `go.mod` file—is essential for **managing dependencies** and keeping your project **organized** and **reliable**.

By using commands like `go mod init`, `go get`, `go mod tidy`, and understanding imports, you can build cleaner, more maintainable Go applications.
