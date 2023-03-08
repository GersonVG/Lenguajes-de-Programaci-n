package main

import "fmt"

// Gira la secuencia a la que apunta arr en n posiciones a la izquierda o a la derecha
// dependiendo del valor de dir (0 = izquierda, 1 = derecha)
func rotarSecuencia(arr *[]int, n int, dir int) {
	length := len(*arr)
	if n == 0 || length == 0 {
		// No se necesita rotación si n es cero o la secuencia está vacía
		return
	}
	// Normalice n para que esté dentro del rango [0, longitud)
	if n < 0 {
		n = (n%length + length) % length
	} else {
		n = n % length
	}
	if dir == 0 {
		// Rotar a la izquierda por n posiciones
		reverseSlice(arr, 0, n-1)
		reverseSlice(arr, n, length-1)
		reverseSlice(arr, 0, length-1)
	} else {
		// Rotar a la derecha n posiciones
		reverseSlice(arr, length-n, length-1)
		reverseSlice(arr, 0, length-n-1)
		reverseSlice(arr, 0, length-1)
	}
}

//Invierta los elementos en el segmento al que apunta arr en el rango [inicio, final]
func reverseSlice(arr *[]int, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func main() {
	// Probar la función de rotarSecuencia
	seq := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Secuencia original:", seq)
	rotarSecuencia(&seq, 3, 0)
	fmt.Println("Secuencia rotada 3 posiciones a la izquierda:", seq)
	rotarSecuencia(&seq, 2, 1)
	fmt.Println("Secuencia rotada 2 posiciones a la derecha:", seq)
}
