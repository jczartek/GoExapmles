package main

// Napisz funkcję reverse(s) odwrcającą kolejność znaków tekstu w argumencie s.
// Zastosuj ją w programie odwracającym kolejno wszystkie wiersze wejściowe.

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	str := []rune(s)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}

	return string(str)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(reverse(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
