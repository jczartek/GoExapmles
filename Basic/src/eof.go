package main

// Napisz program wypisujący wartość i typ EOF.

import (
	"fmt"
	"io"
)

func main() {
	fmt.Printf("Value EOF: %s\nType EOF: %T\n", io.EOF, io.EOF)
}
