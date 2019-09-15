package main

import (
	"fmt"
	"icm.ferth/practica1/v3/racional"
)

func main() {
	var n, d int
	a, b := racional.New(), racional.New()
	fmt.Print("Ingresar numerador y denominardor: ")
	fmt.Scanf("%d %d", &n, &d)
	a.SetNumerador(n)
	a.SetDenominador(d)
	fmt.Print("Ingresar numerador y denominardor: ")
	fmt.Scanf("%d %d", &n, &d)
	b.SetNumerador(n)
	b.SetDenominador(d)
	r := racional.Suma(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, r)
}
