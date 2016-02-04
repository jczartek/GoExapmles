package main

/*
 * Napisz program, który podczas kopiowania wejścia na wyjście zastępuje każdy
 * znak tabulacji przez sekwencję znaków \t, każdy znak \ przez dwa takie znaki.
 */

import (
	"io"
	"os"
)

func main() {
	b := make([]byte, 1)

	_, eof := os.Stdin.Read(b)
	for eof != io.EOF {
		if b[0] == '\t' {
			os.Stdout.Write([]byte("\\t"))
		} else if b[0] == '\\' {
			os.Stdout.Write([]byte("\\\\"))
		} else {
			os.Stdout.Write(b)
		}
		_, eof = os.Stdin.Read(b)
	}
}
