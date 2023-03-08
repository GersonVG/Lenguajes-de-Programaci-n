package main

import "fmt"

func main() {
    var n int
    fmt.Print("Ingrese un número impar mayor o igual a 3: ")
    fmt.Scan(&n)
    
    // Validar que el número ingresado sea impar y mayor o igual a 3
    if n%2 == 0 || n < 3 {
        fmt.Println("Número inválido.")
        return
    }

    // Imprimir el rombo
    for i := 0; i < n/2+1; i++ {
        for j := 0; j < n/2-i; j++ {
            fmt.Print(" ")
        }
        for j := 0; j < 2*i+1; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
    for i := n/2 - 1; i >= 0; i-- {
        for j := 0; j < n/2-i; j++ {
            fmt.Print(" ")
        }
        for j := 0; j < 2*i+1; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

