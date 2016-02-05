package main

/*
 * Napisz program, który każde słowo wejściowe wypisze w osobnym wierszu.
 */

import (
	"io"
	"os"
)

const (
	in int = iota
	out
)

func main() {

	b := make([]byte, 1)
	state := out

	_, eof := os.Stdin.Read(b)

	for eof != io.EOF {
		if b[0] == ' ' || b[0] == '\n' || b[0] == '\t' && state == in {
			os.Stdout.Write([]byte("\n"))
			state = out
		} else if state == out {
			state = in
			os.Stdout.Write(b)
		} else {
			os.Stdout.Write(b)
		}
		_, eof = os.Stdin.Read(b)
	}
}
