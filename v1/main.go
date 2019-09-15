package main

import "fmt"

//Equivalente a una clase en java,
//todos los atributos son publicos a nivel de paquete
type racional struct {
	numerador   int
	denominador int
}

//Equivalente a un constructor vacio en java
func newRacional() *racional {
	return &racional{}
}

//equivalente a un metodo estatico, es global a nivel de paquete
//y no se aplica sobre alguna instancia de alguna clase
func sumaRacionales(a, b *racional) *racional {
	r := newRacional()
	mcm := lcm(a.denominador, b.denominador)
	r.denominador = mcm
	r.numerador = (mcm / a.denominador * a.numerador) + (mcm / b.denominador * b.numerador)
	mcd := gcd(r.numerador, r.denominador)
	r.numerador /= mcd
	r.denominador /= mcd
	return r
}

//maximo comun divisor usando algoritmo de Euclides
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

//minimo comun multiplo lcm(a,b)= (a*b) /gcd(a,b)
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func main() {
	a, b := newRacional(), newRacional()
	fmt.Print("Ingresar numerador y denominardor: ")
	fmt.Scanf("%d %d", &a.numerador, &a.denominador)
	fmt.Print("Ingresar numerador y denominardor: ")
	fmt.Scanf("%d %d", &b.numerador, &b.denominador)
	r := sumaRacionales(a, b)
	fmt.Printf("%d/%d + %d/%d = %d/%d\n", a.numerador, a.denominador, b.numerador, b.denominador, r.numerador, r.denominador)
}
