# Go Fiber Code Service

A simple backend service built with Go Fiber that can run code, auto-fix common mistakes, and provide help responses.

## Features

- **Run Code API**: Execute code and get simulated output
- **Auto-Fix API**: Automatically fix common code issues
- **Help API**: Get help based on keyword matching

## Prerequisites

- Go 1.21 or higher

## Installation

1. Install dependencies:
```bash
go mod download
```

2. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:3000`

## API Endpoints

### 1. Run Code - POST `/run`

Execute code and return output or errors.

**Request Body:**
```json
{
  "code": "console.log('Hello World')"
}
```

**Response:**
```json
{
  "output": "Hello, World!"
}
```

### 2. Auto-Fix Code - POST `/autofix`

Automatically fix common code issues:
- Add missing semicolons
- Fix indentation
- Remove extra spaces
- Correct bracket issues

**Request Body:**
```json
{
  "code": "let x = 10\nconsole.log(x"
}
```

**Response:**
```json
{
  "output": "let x = 10;\nconsole.log(x);"
}
```

### 3. Help - POST `/help`

Get help based on your query using keyword matching.

**Request Body:**
```json
{
  "query": "How do I use loops?"
}
```

**Response:**
```json
{
  "output": "Use 'for' loops for iteration. Syntax: for (init; condition; increment) { code }"
}
```

## Testing

### Option 1: HTML Interface (Recommended)

1. Start the server: `go run main.go`
2. Open your browser and go to: `http://localhost:3000`
3. Use the three sections to test each API

### Option 2: cURL Commands

**Test Run API:**
```bash
curl -X POST http://localhost:3000/run \
  -H "Content-Type: application/json" \
  -d '{"code": "function add(a, b) { return a + b; }"}'
```

**Test Auto-Fix API:**
```bash
curl -X POST http://localhost:3000/autofix \
  -H "Content-Type: application/json" \
  -d '{"code": "let x = 5\nlet y = 10\nconsole.log(x + y"}'
```

**Test Help API:**
```bash
curl -X POST http://localhost:3000/help \
  -H "Content-Type: application/json" \
  -d '{"query": "what is a function?"}'
```

### Option 3: Postman Collection

Import these requests into Postman:

1. Create a new collection named "Go Fiber Code Service"
2. Add three POST requests with the endpoints above
3. Set `Content-Type: application/json` header
4. Use the request bodies from the examples

## Auto-Fix Rules

The auto-fix API applies these fixes:
- Adds missing semicolons to statement lines
- Normalizes indentation based on bracket depth
- Removes extra whitespace
- Adds missing closing brackets
- Adds missing closing parentheses

## Help Topics

The help API recognizes these keywords:
- loop, for, while
- function, parameter, return
- variable, array, object, string, number
- if (conditionals)
- error, debug, syntax
- semicolon, bracket
- class, import

## Project Structure

```
goFiber/
├── main.go           # Main application with all API endpoints
├── go.mod            # Go module dependencies
├── static/
│   └── index.html    # HTML testing interface
|   └── style.css     # For styling
└── README.md         # This file
```
