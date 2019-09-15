package racional

import "strconv"

//Racional es equivalente a una clase en java,
//todos los atributos son privados fuera del paquete racional
type Racional struct {
	numerador   int
	denominador int
}

//New es equivalente a un constructor vacio en java
func New(n, d int) *Racional {
	r := &Racional{numerador: n, denominador: d}
	r.simplificar()
	return r
}

//Equivalente al metodo toString en java
func (r *Racional) String() string {
	str := strconv.Itoa(r.numerador)
	if r.denominador > 1 {
		str += "/" + strconv.Itoa(r.denominador)
	}
	return str
}

//SetNumerador es el metodo setter de numerador
func (r *Racional) SetNumerador(n int) {
	r.numerador = n
	r.simplificar()
}

//SetDenominador es el metodo setter de numerador
func (r *Racional) SetDenominador(d int) {
	r.denominador = d
	if d == 0 {
		r.denominador = 1
	}
	r.simplificar()
}

//simplificar es un metodo privado fuera del paquete
//y que se aplica sobre un racional r
func (r *Racional) simplificar() {
	mcd := gcd(r.numerador, r.denominador)
	r.numerador /= mcd
	r.denominador /= mcd
}

//Suma es equivalente a un metodo estatico, es publico fuera del paquete
//y no se aplica sobre alguna instancia de alguna clase
func Suma(a, b *Racional) *Racional {
	mcm := lcm(a.denominador, b.denominador)
	rd := mcm
	rn := (mcm / a.denominador * a.numerador) + (mcm / b.denominador * b.numerador)
	return New(rn, rd)
}
