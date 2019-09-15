## Practica 1 de APANLOO

Esta practica esta escrita en el lenguaje [Go](https://golang.org/) (Golang). Go es un lenguaje de programación de sistemas semejante al lenguaje C pero con algunas características de los lenguajes de programación modernos, además de ser un lenguaje orientado a objetos aunque esta basado en estructuras y no en clases.

En Go el encapsulamiento se maneja a nivel de paquetes y no de clases como en otros lenguajes orientados a objetos. 

Todos los nombres de variables, constantes, funciones o tipos de datos que comienzan con una letra mayúscula son considerados como públicos y son visibles dentro y fuera del paquete en el que son declarados.

Los nombres que inician con una letra minúscula son considerados como privados y son visibles solo dentro del paquete en el que son declarados. [Ejemplo online](https://gotour-es.appspot.com/#6) 

Para acceder a los elementos públicos de un paquete desde fuera de este es necesario importarlo y se accede a traves del nombre del paquete seguido de un punto y el nombre del elemento.

En Go el siguiente código es equivalente a una clase Hola con 2 atributos en java. [Ejemplo online](https://gotour-es.appspot.com/#25)
``` go
type Hola struct{
    atrib1 int
    atrib2 string
}
```

Los métodos son funciones en las que se especifica un objeto sobre el cual se aplica la llamada al método. Por ejemplo.

``` go
package main

type Hola struct{
    atrib1 int
    atrib2 string
}

func (h Hola) HolaMundo(){
    fmt.Println("Hola",h.atrib2)
}

func main(){
    hola := Hola{atrib2: "Mundo"} 
    hola.HolaMundo()
}

```
En este ejemplo el método HolaMundo() se puede llamar desde cualquier objeto de tipo Hola, internamente el método nombra como ``h`` al objeto que lo llama [ejemplo online](https://gotour-es.appspot.com/#50) , esto es similar al funcionamiento de la palabra reservada ``this`` en java. La salida de este programa es ``Hola Mundo``.
