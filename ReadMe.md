# py-go-bridge
A lightweight framework for seamlessly bridging Python and Go code. Provides type-safe bindings with automatic type conversion and standardized error handling.

## Features
Automatic type conversion between Python and Go
Type validation using Pydantic models
Standardized response format
Simple API design
Header file generation for C bindings

## Quick Start
1. Define your input type in Python:
2. Write your Go function:
3. Use the bridge:

## Response Format

All Go functions return a standardized response:
```json
{
    "status": "success/error",
    "code": 200,  # 200 success, 400 param error, 500 internal error
    "message": "ok",
    "data": {}, # actual return data
    "timing": 100 # execution time in ms
}
```

## Type Conversion

The framework automatically handles conversion between Python and C types:

- str -> char
- int -> int64_t
- float -> double
- bool -> bool
- datetime -> char (ISO format)
- list -> pointer to array
- dict -> char (JSON)

## Development
Build Go library:

cd py_go_bridge/go
go build -buildmode=c-shared -o lib.so

## License
MIT