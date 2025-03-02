# CLI Calculator

A simple command-line calculator built in Go, supporting basic arithmetic operations: Addition, Subtraction, Multiplication, and Division. The CLI also includes commands for clearing the screen, displaying help, and exiting the program.

---

## 📌 Features

- **Addition (****`add`****)**: Add multiple numbers.
- **Subtraction (****`sub`****)**: Subtract one number from another.
- **Multiplication (****`mul`****)**: Multiply multiple numbers.
- **Division (****`div`****)**: Divide two numbers (supports floating-point division and prevents division by zero).
- **Clear (****`clear`****)**: Clears the terminal screen.
- **Help (****`help`****)**: Lists all available commands.
- **Exit (****`exit`****)**: Exits the application.

---

## 🚀 Installation & Setup

### Prerequisites

Ensure you have **Go** installed on your system. You can download it from:
[https://go.dev/dl/](https://go.dev/dl/)

### Clone the Repository

```sh
git clone https://github.com/EhsaasChaudhary/GO_Calculator_CLI
cd cli-calculator
```

### Build the Project

```sh
go build -o calculator
```

### Run the CLI Application

```sh
./calculator
```

---

## 🔹 Usage

### 1️⃣ Help Command

```sh
help
```

Displays all available commands in the following order:

- Help
- Addition
- Subtraction
- Multiplication
- Division
- Clear
- Exit

To get detailed information about a specific command:

```sh
help add
help sub
```

### 2️⃣ Addition (`add`)

```sh
add <num1> <num2> ... <numN>
```

Example:

```sh
add 2 5 8 8
```

**Output:** `2 + 5 + 8 + 8 = 23`

### 3️⃣ Subtraction (`sub`)

```sh
sub <num1> <num2>
```

Example:

```sh
sub 2 1
```

**Output:** `2 - 1 = 1`

Supports negative results:

```sh
sub 1 2
```

**Output:** `1 - 2 = -1`

### 4️⃣ Multiplication (`mul`)

```sh
mul <num1> <num2> ... <numN>
```

Example:

```sh
mul 2 5 8 8
```

**Output:** `2 * 5 * 8 * 8 = 640`

### 5️⃣ Division (`div`)

```sh
div <num1> <num2>
```

Example:

```sh
div 10 5
```

**Output:** `10 / 5 = 2.0`

Supports floating-point results:

```sh
div 1 2
```

**Output:** `1 / 2 = 0.5`

Handles division by zero:

```sh
div 10 0
```

**Output:** `❌ Error: Division by zero is not allowed.`

### 6️⃣ Clear (`clear`)

```sh
clear
```

Clears the terminal screen.

### 7️⃣ Exit (`exit`)

```sh
exit
```

Closes the application.

---

## 📜 License

This project is licensed under the MIT License.

---

### 🎯 Author

**Your Name**\
GitHub: [EhsaasChaudhary](https://github.com/EhsaasChaudhary)

