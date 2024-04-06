# Simple Go PrintRune Function

This is a simple Go program that prints a rune to the writer.

Made for the purpose of learning Go.

## Usage

```go
package main

import (
    "github.com/wtfuk/s01"
)

func main() {
    s01.PrintRune('🚀')
}
```

## Testing

```sh
go test -race -v
```

## Makefile Use

- To build your project, run `make build`.
- To install dependencies, run `make deps`.
- To update dependencies, run `make update`.
- To run tests, run `make test`.
- To clean up the project directory, run `make clean`.
- To run your project, use `make run`.
- To cross-compile for a specific OS, run one of the cross-compilation targets, e.g., `make build-linux`.

## License

SSALv2

## Author

- [Sagar Yadav](https://linkedin.com/in/sagaryadav)
