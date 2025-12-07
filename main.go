package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type CodeRequest struct {
	Code string `json:"code"`
}

type QueryRequest struct {
	Query string `json:"query"`
}

type Response struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/run", runCode)
	app.Post("/autofix", autoFix)
	app.Post("/help", help)

	app.Static("/", "./static")

	fmt.Println("Server running on http://localhost:3000")
	app.Listen(":3000")
}

func runCode(c *fiber.Ctx) error {
	var req CodeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(Response{Error: "Invalid request body"})
	}

	if req.Code == "" {
		return c.Status(400).JSON(Response{Error: "Code is required"})
	}

	output := simulateExecution(req.Code)
	return c.JSON(Response{Output: output})
}

func simulateExecution(code string) string {
	code = strings.TrimSpace(code)

	if strings.Contains(code, "syntax error") {
		return "Error: syntax error in code"
	}
	if strings.Contains(code, "undefined") {
		return "Error: undefined variable or function"
	}
	if strings.Contains(code, "null pointer") {
		return "Error: null pointer exception"
	}

	if strings.HasPrefix(code, "print") || strings.Contains(code, "console.log") {
		return "Hello, World!"
	}
	if strings.Contains(code, "function") || strings.Contains(code, "def ") {
		return "Function defined successfully"
	}
	if strings.Contains(code, "for") || strings.Contains(code, "while") {
		return "Loop executed successfully"
	}

	return "Code executed successfully"
}

func autoFix(c *fiber.Ctx) error {
	var req CodeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(Response{Error: "Invalid request body"})
	}

	if req.Code == "" {
		return c.Status(400).JSON(Response{Error: "Code is required"})
	}

	fixedCode := applyAutoFixes(req.Code)
	return c.JSON(Response{Output: fixedCode})
}

func applyAutoFixes(code string) string {
	lines := strings.Split(code, "\n")
	var fixedLines []string

	for _, line := range lines {
		fixed := line

		fixed = strings.TrimRight(fixed, " \t")

		fixed = regexp.MustCompile(`\s+`).ReplaceAllString(fixed, " ")

		if strings.TrimSpace(fixed) != "" && !strings.HasSuffix(strings.TrimSpace(fixed), ";") &&
			!strings.HasSuffix(strings.TrimSpace(fixed), "{") &&
			!strings.HasSuffix(strings.TrimSpace(fixed), "}") &&
			!strings.HasPrefix(strings.TrimSpace(fixed), "//") &&
			!strings.HasPrefix(strings.TrimSpace(fixed), "#") &&
			!strings.Contains(fixed, "if ") &&
			!strings.Contains(fixed, "for ") &&
			!strings.Contains(fixed, "while ") &&
			!strings.Contains(fixed, "function") &&
			!strings.Contains(fixed, "def ") &&
			!strings.Contains(fixed, "class ") {
			if !strings.HasSuffix(strings.TrimSpace(fixed), ";") {
				fixed = strings.TrimRight(fixed, " ") + ";"
			}
		}

		openBrackets := strings.Count(fixed, "{")
		closeBrackets := strings.Count(fixed, "}")
		if openBrackets > closeBrackets {
			fixed += " }"
		}

		openParens := strings.Count(fixed, "(")
		closeParens := strings.Count(fixed, ")")
		if openParens > closeParens {
			fixed += ")"
		}

		fixedLines = append(fixedLines, fixed)
	}

	result := strings.Join(fixedLines, "\n")

	result = normalizeIndentation(result)

	return result
}

func normalizeIndentation(code string) string {
	lines := strings.Split(code, "\n")
	var normalized []string
	indentLevel := 0

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			normalized = append(normalized, "")
			continue
		}

		if strings.Contains(trimmed, "}") || strings.Contains(trimmed, "]") {
			if indentLevel > 0 {
				indentLevel--
			}
		}

		indent := strings.Repeat("  ", indentLevel)
		normalized = append(normalized, indent+trimmed)

		if strings.Contains(trimmed, "{") || strings.Contains(trimmed, "[") {
			indentLevel++
		}
	}

	return strings.Join(normalized, "\n")
}

func help(c *fiber.Ctx) error {
	var req QueryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(Response{Error: "Invalid request body"})
	}

	if req.Query == "" {
		return c.Status(400).JSON(Response{Error: "Query is required"})
	}

	helpText := getHelpResponse(req.Query)
	return c.JSON(Response{Output: helpText})
}

func getHelpResponse(query string) string {
	query = strings.ToLower(query)

	keywords := map[string]string{
		"loop":      "Use 'for' loops for iteration. Syntax: for (init; condition; increment) { code }",
		"for":       "For loops: for (let i = 0; i < 10; i++) { console.log(i); }",
		"while":     "While loops: while (condition) { code }. Make sure condition eventually becomes false.",
		"function":  "Functions are reusable code blocks. Syntax: function name(params) { return value; }",
		"variable":  "Declare variables using let or const. Example: let x = 10; const PI = 3.14;",
		"array":     "Arrays store multiple values. Example: let arr = [1, 2, 3]; Access: arr[0]",
		"object":    "Objects store key-value pairs. Example: let obj = {name: 'John', age: 30};",
		"if":        "Conditional statements: if (condition) { code } else { alternative }",
		"error":     "Check for syntax errors, missing brackets, or undefined variables. Use console.log for debugging.",
		"debug":     "Use console.log() to print values and trace execution. Check browser console for errors.",
		"syntax":    "Common syntax errors: missing semicolons, unmatched brackets, typos in keywords.",
		"semicolon": "Semicolons end statements in many languages. Example: let x = 5;",
		"bracket":   "Match opening and closing brackets: (), {}, []. Each opening needs a closing.",
		"string":    "Strings are text in quotes. Example: let str = 'Hello'; or let str = \"World\";",
		"number":    "Numbers don't need quotes. Example: let num = 42; let pi = 3.14;",
		"return":    "Use return to output a value from a function. Example: return result;",
		"parameter": "Parameters are function inputs. Example: function add(a, b) { return a + b; }",
		"class":     "Classes define objects. Syntax: class Name { constructor() {} method() {} }",
		"import":    "Import modules: import { name } from 'module'; or const name = require('module');",
	}

	for keyword, response := range keywords {
		if strings.Contains(query, keyword) {
			return response
		}
	}

	return "I can help with: loops, functions, variables, arrays, objects, conditionals, debugging, and common syntax issues. Try asking about a specific topic."
}
