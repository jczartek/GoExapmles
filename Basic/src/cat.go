package main

// Napisz program przepisujący wejście na wyjście

import (
	"bufio"
	"io"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	c, eof := reader.ReadByte()

	for eof != io.EOF {
		writer.WriteByte(c)
		c, eof = reader.ReadByte()
	}

	writer.Flush()
}
