package racional

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
