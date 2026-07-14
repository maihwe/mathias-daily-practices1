# Ascii Art Web

## Description
Ascii-art-web is a web-based graphical user interface (GUI) version of the ASCII Art generator project. It allows users to input text via a web browser form, select from three different banner styles (Standard, Shadow, and Thinkertoy), and see the converted text instantly rendered on the page.

## Authors
* [MATHIAS IHWE (maihwe)]

## Usage: how to run
To run this web application locally, ensure you have Go installed on your machine. 

1. Download or clone this project directory.
2. Place the banner files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`) in the root folder.
3. Open your terminal in the root directory and run:
   ```bash
   go run .
   ```
4. Open your web browser and navigate to: `http://localhost:8081`

## Implementation details: algorithm
The project is built entirely using Go's standard library packages without any external frameworks. 

* **Routing & Server:** The `net/http` package handles requests. A `GET /` request serves the home template, while a `POST /ascii-art` route processes incoming user forms, validating form inputs and ensuring correct HTTP Status Codes (200, 400, 404, 500).
* **Templates:** The `html/template` package parses `index.html` at startup and uses context-safe pipeline binding (`{{.}}`) inside standard `<pre>` HTML tags to keep character alignment intact.
* **ASCII Art Algorithm:** The backend loads the selected banner file as a slice of strings (`[]string`). Since every character block consists of 8 art lines and a trailing separator newline, the algorithm loops across 8 rows horizontally, indexing characters dynamically using the formula `(asciiValue - 32) * 9 + 1` to prevent vertical text stacking.
