package main

/*
 * Napisz program wypisujący wszystkie wiersze wejściowe dłuższe niż n znaków.
 */

import (
	"bufio"
	"fmt"
	"os"
)

const n int = 70

func main() {
	var line string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if line = scanner.Text(); len(line) > n {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
