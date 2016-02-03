package main

// Spróbuj sprawdzić co się stanie, gdy w argumencie funkcji Printf wystąpi
// sekwencja specjalna \c, w której c jest dowolnym znakiem od wyżej wymienionych.

import "fmt"

func main() {
  fmt.Printf("Hello World\y")
  fmt.Printf("Hello World\7")
  fmt.Printf("Hello World\?")
}

// Program się nie skompiluje. Kompilator Go wyświetla błąd: nieznane sekwencje
// specjalne.
