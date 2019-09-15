package racional

import "strconv"

//Racional es equivalente a una clase en java,
//todos los atributos son publicos fuera del paquete racional
type Racional struct {
	Numerador   int
	Denominador int
}

//New es equivalente a un constructor vacio en java
func New() *Racional {
	return &Racional{Denominador: 1}
}

//Equivalente al metodo toString en java
func (r *Racional) String() string {
	str := strconv.Itoa(r.Numerador)
	if r.Denominador != 1 {
		str += "/" + strconv.Itoa(r.Denominador)
	}
	return str
}

//Suma es equivalente a un metodo estatico, es publico fuera del paquete racional
//y no se aplica sobre alguna instancia de alguna clase
func Suma(a, b *Racional) *Racional {
	r := New()
	mcm := lcm(a.Denominador, b.Denominador)
	r.Denominador = mcm
	r.Numerador = (mcm / a.Denominador * a.Numerador) + (mcm / b.Denominador * b.Numerador)
	mcd := gcd(r.Numerador, r.Denominador)
	r.Numerador /= mcd
	r.Denominador /= mcd
	return r
}
