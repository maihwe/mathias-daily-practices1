
# CodeCrafters — Base Converter (CLI Tool)

## Overview

The **Base Converter** is a command-line tool written in Go that converts numbers between different bases:

* **Hexadecimal → Decimal**
* **Binary → Decimal**
* **Decimal → Binary & Hexadecimal**

This project reinforces understanding of:

* Number systems (base 2, 10, 16)
* String parsing and validation
* CLI interaction using loops and user input
* Error handling with Go's `strconv` package

---

## Features

* Supports three bases:

  * `hex` (base 16)
  * `bin` (base 2)
  * `dec` (base 10)

* Conversion behavior:

  * `hex` → outputs **decimal**
  * `bin` → outputs **decimal**
  * `dec` → outputs **binary and hex**

* Continuous execution until user exits

* Handles invalid inputs gracefully

* Supports negative decimal numbers

* Outputs hexadecimal in **uppercase**


## How to Run

### 1. Navigate to project directory

```bash
cd thecodecrafterthon-day2
```

### 2. Run the program

```bash
go run main.go
``

## Usage

The program expects input in this format:

```bash
convert <value> <base>
```

### Examples

```bash
> convert 1E hex
✦ Decimal: 30

> convert 10 bin
✦ Decimal: 2

> convert 255 dec
✦ Binary:  11111111
✦ Hex:     FF
```

---

## Exit

To stop the program:

```bash
quit
```

---

## Input Validation

The program handles invalid inputs safely:

| Input       | Reason                  |
| ----------- | ----------------------- |
| `1G hex`    | Invalid hex character   |
| `10201 bin` | Invalid binary digit    |
| `abc dec`   | Not a valid decimal     |
| empty input | Ignored / error message |

---

## Key Concepts Used

### 1. `strconv.ParseInt`

Used to convert strings into integers with a specified base:

```go
strconv.ParseInt(value, base, 64)
```

* Automatically validates input
* Returns an error if input is invalid

---

### 2. Formatting Output

```go
fmt.Printf("%b", n) // Binary
fmt.Printf("%X", n) // Hex (uppercase)
```

---

### 3. CLI Loop

The program continuously reads user input:

```go
for {
    scanner.Scan()
    ...
}
```


## Project Structure

```
thecodecrafterthon-day2/
├── main.go
└── README.md
```


## 👤 Author

* Name: [Otete Benjamin Agogo]
* Project: CodeCrafters — Base Converter
* Language: Go (Golang)

---
