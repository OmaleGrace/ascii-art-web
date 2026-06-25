# ASCII Art Web application built with Go and HTML templates. Supports multiple banner styles, input validation, and dynamic style switching.

## Description

ASCII-Art-Web is a web application written in Go that converts user input into ASCII art using different banner styles. Users can enter text through a web interface, generate ASCII art, and then choose an alternative font style without retyping their text.

The application supports multiple banner files such as **standard**, **shadow**, and **thinkertoy**, allowing users to customize the appearance of their generated ASCII art.

---

## Authors

* Omale Grace

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
* Instant style comparison without regenerating text manually.

---

## Project Structure

```
## Project Structure

ascii-art-web/
├── ascii/
│   ├── generate.go
│   ├── loadbanner.go
│   └── render.go
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── templates/
│   ├── input.html
│   ├── output.html
│   └── base.html
├── ArtHandler.go
├── PageHandler.go
├── SwitchHandler.go
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

1. Click one of the available banner style links:

* standard
* shadow
* thinkertoy

2. The application automatically regenerates the ASCII art using the selected banner.
3. The original text is preserved during the switch.
4. The page refreshes with the newly styled output.

This allows users to quickly compare different ASCII-art representations without re-entering text or submitting a new form.

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

| Method | Route             | Description                                            |
| ------ | ------------------|--------------------------------------------------------|
| GET    | /                 | Displays the home page                                 |
| POST   | /ascii-art        | Generates ASCII art                                    |
| POST   | /ascii-art-switch | Regenerates output with using a different banner style |

---


---

### 3. Error Handling section


## Error Handling

The application handles:

- **400 Bad Request**
  - Empty input
  - Missing query parameters
  - Invalid requests

- **404 Not Found**
  - Unknown routes

- **405 Method Not Allowed**
  - Unsupported HTTP methods

- **500 Internal Server Error**
  - Banner loading failures
  - Template parsing or execution failures

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

The application supports dynamic banner switching from the output page.


1. The generated output page contains links for all available banner styles.
2. When a user clicks a banner style link, the original text and selected banner are sent as query parameters.
3. The server loads the requested banner file.
4. The ASCII art is regenerated using the same text.
5. The updated result is displayed immediately.

This feature enables fast comparison between different banner styles while preserving the user's original input.

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
