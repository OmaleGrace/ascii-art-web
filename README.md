# ASCII-Art-Web

## Description

ASCII-Art-Web is a web application written in Go that converts user input into ASCII art using different banner styles. Users can enter text through a web interface, generate ASCII art, and then choose an alternative font style without retyping their text.

The application supports multiple banner files such as **standard**, **shadow**, and **thinkertoy**, allowing users to customize the appearance of their generated ASCII art.

---

## Authors

* Ooja Omale

---

## Features

* Generate ASCII art from user input.
* Support for multiple banner styles:

  * Standard
  * Shadow
  * Thinkertoy
* Web-based user interface.
* Proper HTTP status handling.
* Input validation and error handling.
* Dynamic banner switching from the output page.
* Preservation of user input when changing styles.

---

## Project Structure

```
ascii-art-web/
├── ascii/
│   ├── generate.go
│   ├── LoadBanner.go
│   └── render.go
├── banner/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── templates/
│   ├── index.html
│   └── output.html
├── ArtHandler.go
├── PageHandler.go
├── main.go
└── go.mod
```

---

## Usage

### Clone the Repository

```bash
git clone https://github.com/omalegrace2009-g/ascii-art-web.git
cd ascii-art-web
```

### Run the Application

```bash
go run .
```

The server starts on:

```
http://localhost:8080
```

---

## How to Use

### Generate ASCII Art

1. Open the application in your browser.
2. Enter the desired text.
3. Select a banner style.
4. Click **Generate**.
5. The generated ASCII art will be displayed.

### Change the Banner Style

After the ASCII art is generated:

1. Use the banner selection form available on the output page.
2. Choose a different style.
3. Click **Change Style**.
4. The application regenerates the same text using the selected banner.

This allows users to compare different ASCII-art representations without re-entering their text.

---

## Supported Banner Styles

### Standard

A clean and traditional ASCII-art font.

### Shadow

Produces a bold shadow-like appearance.

### Thinkertoy

Creates a lightweight and decorative representation.

---

## HTTP Endpoints

| Method | Route         | Description                            |
| ------ | ------------- | -------------------------------------- |
| GET    | /             | Displays the home page                 |
| POST   | /ascii-art    | Generates ASCII art                    |
| POST   | /change-style | Regenerates output with another banner |

---

## Error Handling

The application handles:

* 400 Bad Request

  * Empty input
  * Invalid banner selection
  * Malformed requests

* 404 Not Found

  * Unknown routes

* 405 Method Not Allowed

  * Unsupported HTTP methods

* 500 Internal Server Error

  * Banner loading failures
  * Template execution errors
  * Unexpected server-side issues

---

## Implementation Details

### Banner Loading

Banner files are read and parsed into memory. Each printable ASCII character is mapped to its corresponding 8-line representation.

### Rendering Algorithm

1. Receive user input.
2. Split input into lines.
3. Load the selected banner.
4. Convert each character to its ASCII-art equivalent.
5. Concatenate the resulting lines.
6. Return the generated output to the template.

### Banner Switching

When a user selects a new banner on the output page:

1. The original text is preserved using hidden form fields.
2. The newly selected banner is submitted.
3. The server regenerates the ASCII art.
4. The output page refreshes with the updated style.

---

## Example

Input:

```
Hello
```

Using the Standard banner:

```
 _   _      _ _
| | | | ___| | | ___
| |_| |/ _ \ | |/ _ \
|  _  |  __/ | | (_) |
|_| |_|\___|_|_|\___/
```

Changing the style to Shadow or Thinkertoy will regenerate the same text using their respective banner representations.

---

## License

This project was developed for educational purposes as part of the ASCII-Art-Web project.
