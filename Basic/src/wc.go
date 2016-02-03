package main

// Napisz program zliczający znaki odstępu, tabulacji i nowego wiersza.

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var nb, nt, nl int

	reader := bufio.NewReader(os.Stdin)
	c, err := reader.ReadByte()

	for err != io.EOF {
		if c == ' ' {
			nb++
		} else if c == '\t' {
			nt++
		} else if c == '\n' {
			nl++
		}
		c, err = reader.ReadByte()
	}
	fmt.Println(nb, nt, nl)
}
