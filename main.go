package main

import "fmt"

type Correos struct {
	nombre string
}
type listaC [3]Correos

type Persona struct {
	nombre       string
	listaCorreos *listaC //listaCorreos listaC
}

var listaPersonas []Persona

func main() {
	var lista listaC
	lista[0] = Correos{"correo1"}
	lista[1] = Correos{"correo2"}
	lista[2] = Correos{"correo3"}

	/*Julio := Persona{"Julio", &lista}
	Cesar := Persona{"Cesar", &lista}
	Jose := Persona{"Jose", &lista}
	Arnoldo := Persona{"Arnoldo", &lista} */

	Julio := Persona{"Julio", lista}
	Cesar := Persona{"Cesar", lista}
	Jose := Persona{"Jose", lista}
	Arnoldo := Persona{"Arnoldo", lista}

	listaPersonas := append(listaPersonas, Julio)
	listaPersonas = append(listaPersonas, Cesar)
	listaPersonas = append(listaPersonas, Jose)
	listaPersonas = append(listaPersonas, Arnoldo)

	imprimirListaPersona(listaPersonas) //fmt.Println(listaPersonas)
	per := buscarPersona(&listaPersonas, "Julio")

	println(per)
	/* imprimirPersona(*per)
	println() */
	println(&listaPersonas[0])
	/* imprimirPersona(listaPersonas[0])
	println() */
	//cambiarCorreo(0, &((*per).listaCorreos))
	cambiarCorreo(0, &(listaPersonas[0].listaCorreos))
	//cambiarCorreo(0, listaPersonas[0].listaCorreos)

	fmt.Println(listaPersonas)
	//imprimirListaPersona(listaPersonas)
}

func buscarPersona(listaP *[]Persona, nombre string) *Persona {

	var i int
	for i = 0; i < len(*listaP); i++ {
		if (*listaP)[i].nombre == nombre {
			return &(*listaP)[i]
		}
	}
	return nil
}

func cambiarCorreo(pos int, lcorreos *listaC) {
	(*lcorreos)[pos].nombre = "otro"
}

func imprimirPersona(p Persona) {
	fmt.Print("{")
	fmt.Print(p.nombre + " ")
	i := 0
	for i < 3 {
		fmt.Print(p.listaCorreos[i])
		i++
	}
	fmt.Print("} ")
}

func imprimirListaPersona(l []Persona) {
	fmt.Print("[ ")
	for _, persona := range l {
		imprimirPersona(persona)
	}
	fmt.Println("]")
}
