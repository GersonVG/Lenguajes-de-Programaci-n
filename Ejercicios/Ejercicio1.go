Ejercicio 1 :

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Ingrese su texto: ") //El input para ingresar los caracteres
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()
	numCaracteres := len(texto)                 //Contar el número de caracteres en el texto y función
	numPalabras := len(strings.Fields(texto))   //Sirve para contar el número de palabras en el texto
	numLineas := strings.Count(texto, "\n") + 1 //Acá se cuentan el número de líneas en el texto

	fmt.Println("Número de caracteres:", numCaracteres)
	fmt.Println("Número de palabras:", numPalabras)
	fmt.Println("Número de líneas:", numLineas)
}
