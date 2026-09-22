package main

// print  -> printea literal
// printf -> printea con formato

// el compilador de go es retrolazo q rompe las bolas si importas cosas q no usas

import (
	"fmt" // sirve para imprimir en consola
	"net"
	"os" // sirve para interactuar con el sistema operativo
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
		fmt.Println("Hola, soy el HOST, espero q se conecte algun alma.")
		listener, err := net.Listen("tcp", ":8080") //escucha el puerto 8080
		if err != nil {
			fmt.Println("Error al abrir el puerto:", err) //si hay errpr lo imprime
			return
		}
		fmt.Println("Puerto 8080 abierto, Esperando...")

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexion:", err)
			return
		}
		fmt.Println("Alguien se conecto")

		conn.Close() // se cierra la conexion

	} else if rol == "guest" {
		// se conecta al host
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			fmt.Println("Error al conectar con el host:", err) // en caso de no poder conectarse
			return
		}
		fmt.Println("¡Conexion exitosaa :0 ")

		// se lee la respuesta del host
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error al leer el mensaje:", err)
			return
		}
		// se imprime la respuesta del host
		fmt.Printf("El host dice: %s\n", string(buffer[:n]))
		conn.Close() //cierra la conexion

	} else {
		fmt.Println("Rol incorrecto es o HOST o GUEST")
	}
}
