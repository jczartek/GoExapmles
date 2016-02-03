package main

// Napisz program, który wypisze zestawienie temperatur w skali Fahrenheita
// oraz ich odpowiedniki w skali Celsjusza, wyliczonych według wzoru
// C = (5/9)(F-32). Uwzględnij również nagłówek zestawienia temperatur.

import "fmt"

const (
	lower = 0
	upper = 300
	step  = 20
)

func main() {
	var fahr int

	fmt.Printf("%3s %6s\n", "Fahr", "Cels")
	for fahr = lower; fahr <= upper; fahr += step {
		fmt.Printf("%4d %6.1f\n", fahr, (5.0/9.0)*(float32(fahr)-32.0))
	}
}
