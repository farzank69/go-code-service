#!/bin/bash

echo "=== Testing Go Fiber Code Service ==="
echo ""

echo "1. Testing Run Code API"
echo "Request: POST /run"
curl -X POST http://localhost:3000/run \
  -H "Content-Type: application/json" \
  -d '{"code": "console.log(\"Hello World\")"}' \
  -w "\n"
echo ""

echo "2. Testing Run Code with Function"
curl -X POST http://localhost:3000/run \
  -H "Content-Type: application/json" \
  -d '{"code": "function add(a, b) { return a + b; }"}' \
  -w "\n"
echo ""

echo "3. Testing Auto-Fix API"
echo "Request: POST /autofix"
curl -X POST http://localhost:3000/autofix \
  -H "Content-Type: application/json" \
  -d '{"code": "let x = 10\nlet y = 20\nconsole.log(x + y"}' \
  -w "\n"
echo ""

echo "4. Testing Auto-Fix with Indentation"
curl -X POST http://localhost:3000/autofix \
  -H "Content-Type: application/json" \
  -d '{"code": "if (true) {\nconsole.log(\"test\"\n}"}' \
  -w "\n"
echo ""

echo "5. Testing Help API"
echo "Request: POST /help"
curl -X POST http://localhost:3000/help \
  -H "Content-Type: application/json" \
  -d '{"query": "How do I use loops?"}' \
  -w "\n"
echo ""

echo "6. Testing Help with Different Query"
curl -X POST http://localhost:3000/help \
  -H "Content-Type: application/json" \
  -d '{"query": "what is a function?"}' \
  -w "\n"
echo ""

echo "7. Testing Help with Array Query"
curl -X POST http://localhost:3000/help \
  -H "Content-Type: application/json" \
  -d '{"query": "how to use arrays?"}' \
  -w "\n"
echo ""

echo "=== All tests completed ==="
