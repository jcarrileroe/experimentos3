// Módulo de mi primer programa en go no solo saluda sino que lo traduce a Kirgon
package main

import (
	"fmt"

	"github.com/jcarrileroe/experimentos3/mod1"
	"github.com/jcarrileroe/experimentos3/mod2"
)

func main() {
	saludo := "Hola BME"
	fmt.Println(saludo)
	fmt.Println(mod2.Translate(saludo))
	fmt.Println(mod1.Sumar(2, 3))
}
