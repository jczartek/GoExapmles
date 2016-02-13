package main

// Napisz program znajdujący liczby pierwsze z zakresu od 2 do 100. Wykorzystaj
// do tego metodę zwaną "sitem Eratostenesa"

import "fmt"

func sieveE(n int) []bool {
	primes := make([]bool, n)

	for i := 2; i*i <= n; i++ {
		if primes[i] == false {
			for j := i * i; j < n; j += i {
				primes[j] = true
			}
		}
	}
	return primes
}

func printSieve(p []bool) {
	for i := 2; i < len(p); i++ {
		if p[i] == false {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func main() {
	p := sieveE(100)
	printSieve(p)
}
