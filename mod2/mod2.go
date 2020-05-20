// Este modulo traduce cadenas de texto al complicado lenguage Kirgon
package mod2

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println("Modulo 2: activado")
	rand.Seed(time.Now().Unix())
}

// Traduce cualquier cadena a Kyrgon
// Translate(":Hola")
// devuelve el famoso saludo Kyrgon
func Translate(cadena string) string {
	runes := []rune(cadena)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}
