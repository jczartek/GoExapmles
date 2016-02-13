package main

/*
 * Napisz program obliczający medianę.
 *
 * Medianą uporządkowanego rosnąco ciągu n danych liczbowych a1 <= a2 ... <= an
 * jest:
 * -- dla n nieparzystego:  środkowy wyraz ciągu
 * -- dla n parzystego:     średnia arytmetyczna środkowych wyrazów ciągu
 */

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	var n int
	var numbers []int

	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}
		numbers = append(numbers, n)
	}

	if len(numbers) < 2 {
		log.Print("You must enter at least two numbers!")
		os.Exit(1)
	}

	sort.Ints(numbers)

	switch len(numbers) % 2 {
	case 0:
		middle := len(numbers) / 2
		median := (float32(numbers[middle-1] + numbers[middle])) / 2
		fmt.Printf("Median is %.1f\n", median)
	case 1:
		fmt.Printf("Median is %d\n", numbers[len(numbers)/2])
	}
}
