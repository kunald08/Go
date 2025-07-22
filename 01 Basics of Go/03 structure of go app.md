
# 🏗️ Structure of a Go Application

Proper file and folder structure is essential for building **scalable**, **maintainable**, and **clean** Go applications.

---

## 🧱 Basic Go Project Structure

Ideal for **small projects** and quick prototypes.

### 📁 Example Structure

```
/my-go-app
├── main.go
├── utils.go
├── go.mod
├── go.sum
└── README.md
```

### 🔍 Component Explanation

- **main.go**: Contains the `main()` function; entry point of the application.
- **utils.go**: Utility functions like logging or math helpers; part of the main package.
- **go.mod / go.sum**: Manage project dependencies and versions.
- **README.md**: (Optional) Overview and usage instructions.

### ✅ Benefits

- Simple and minimalistic
- Easy to navigate for new developers
- Perfect for learning, demos, and small apps

---

## 🧩 Complex Go Project Structure (Larger Applications)

Used for **enterprise-grade** or **multi-component** applications.

### 📁 Example Structure

```
/my-large-app
├── cmd/
│   └── myapp/main.go
├── pkg/
│   └── util/util.go
├── internal/
│   ├── config/config.go
│   └── handler/handler.go
├── api/
│   └── v1/api.go
├── scripts/
│   └── setup.sh
├── configs/
│   └── constants.go
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```

### 🔍 Folder Breakdown

- **/cmd**: Application entry points (e.g., `main.go` for different executables)
- **/pkg**: Public packages meant to be imported across the project
- **/internal**: Private logic only used within the project
- **/api**: Versioned API implementation (e.g., REST handlers)
- **/scripts**: Shell scripts for setup, CI/CD, or automation
- **/configs**: Configuration files and environment variables
- **.gitignore**: Files/directories Git should ignore (e.g., `.env`, binaries)
- **go.mod / go.sum**: Dependency management
- **README.md**: Project overview, setup, and usage guide

### ✅ Benefits

- Modular and maintainable
- Clear separation of concerns
- Scales well with project size

---

## 🏗️ Layered Architecture (MVC Example)

Helps manage **complex logic** using logical separation.

### 📁 Example Layout

```
/my-layered-app
├── models/
│   └── user.go
├── controllers/
│   └── user_controller.go
├── config/
│   └── db.go
├── public/
│   ├── styles.css
│   └── logo.png
```

### 🔍 Layer Responsibilities

- **Models**: Define data structures and interact with the database.
- **Controllers**: Handle logic between models and views.
- **Config**: Store settings like DB connections or environment config.
- **Public**: Static files such as HTML, CSS, JS, and images.

### ✅ Benefits

- Promotes clean code separation
- Improves readability and testing
- Ideal for **API-first** or **full-stack** Go apps

---

## 🧾 Conclusion

- Use **Basic Structure** for small, fast projects or learning purposes.
- Use **Complex Structure** for medium to large applications.
- Adopt **Layered Architecture** for scalable and maintainable enterprise-grade systems.
