package main

// print  -> printea literal
// printf -> printea con formato

// el compilador de go es retrolazo q rompe las bolas si importas cosas q no usas

import (
	"fmt" // sirve para imprimir en consola
	"os"  // sirve para interactuar con el sistema operativo
	//"net"
)

func main() {
	//chekeo
	if len(os.Args) < 2 {
		fmt.Println("Error: escribe host o guest ")
		return
	}
	// rol: elegir entre host o guest
	rol := os.Args[1]

	if rol == "host" {
		fmt.Println("Hola, soy el HOST,espero q se conecte algun alma.")

	} else if rol == "guest" {
		fmt.Println("Hola, soy el GUEST. Voy a intentar llamar al host.")

	} else {
		fmt.Println("Rol incorrecto es o HOST o GUEST")
	}
}
