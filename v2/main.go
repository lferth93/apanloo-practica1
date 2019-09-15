package main

import (
	"fmt"
	"icm.ferth/practica1/v2/racional"
)

func main() {
	a, b := racional.New(), racional.New()
	fmt.Print("Ingresar numerador y denominardor: ")
	fmt.Scanf("%d %d", &a.Numerador, &a.Denominador)

	fmt.Print("Ingresar numerador y denominardor: ")
	fmt.Scanf("%d %d", &b.Numerador, &b.Denominador)
	r := racional.Suma(a, b)
	fmt.Printf("%v + %v = %v\n", a, b, r)
}
