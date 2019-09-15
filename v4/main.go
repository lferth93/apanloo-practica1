package main

import (
	"fmt"

	"icm.ferth/practica1/v4/racional"
)

func main() {
	var n, d int
	fmt.Print("Ingresar numerador y denominardor: ")
	fmt.Scanf("%d %d", &n, &d)
	a := racional.New(n, d)
	fmt.Print("Ingresar numerador y denominardor: ")
	fmt.Scanf("%d %d", &n, &d)
	b := racional.New(n, d)
	r := racional.Suma(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, r)
}
