package main

import( 
	"fmt" // para imprimir en consola 
	//"os"  // acceder a entrada del sistema 
)

func main(){
	mensaje := ""

	fmt.Print("Escribe tu mensaje: ")

	fmt.Scanln(&mensaje)

	fmt.Printf("Mensaje capturado %s\n", mensaje)

}