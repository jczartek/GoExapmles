package main

// Napisz program wypisujący zestawienie temperatur Fahrenheita-Celsjusza tak,
// aby wypisywał zesstawienie w odwrotnej kolejności, to znaczy od 300 stopni
// do zera.

import "fmt"

func main() {
	var fahr int

	for fahr = 300; fahr >= 0; fahr -= 15 {
		fmt.Printf("%3d %6.1f\n", fahr, (5.0/9.0)*(float32(fahr)-32.0))
	}
}
