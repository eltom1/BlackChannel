package main

// print  -> printea literal
// printf -> printea con formato

// el compilador de go es retrolazo q rompe las bolas si importas cosas q no usas

import (
	"fmt" // sirve para imprimir en consola
	"net" //sirve para crear conexiones de red
	// sirve para interactuar con el sistema operativo
)

//os.Args[0] -> nombre del programa
//os.Args[1] -> parametro para el host o guest
//os.Args[2] -> parametro para el puerto

func main() {
	//chekeo si no tiene menos de dos parametros tira error

	/*
		if len(os.Args) < 2 {
			fmt.Println("Error: escribe host o guest ")
			return
		}
	*/

	tipo := ""
	puerto := ""

	fmt.Print("ingrese el puerto: ")
	fmt.Scanln(&puerto)

	fmt.Print("ingrese el tipo de usuario(host o guest): ")
	fmt.Scanln(&tipo)

	// rol: elegir entre host o guest
	rol := tipo

	if rol == "host" {
		fmt.Println("Hola, soy el HOST, espero q se conecte algun alma.")
		listener, err := net.Listen("tcp", ":"+puerto) //escucha el puerto 8080

		if err != nil {
			fmt.Println("Error al abrir el puerto:", err) //si hay errpr lo imprime
			return
		}
		fmt.Println("Puerto ", puerto, " abierto, Esperando...")

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexion:", err)
			return
		}
		fmt.Println("Alguien se conecto")

		conn.Write([]byte("se conecto el host \n"))

		conn.Close() // se cierra la conexion

	} else if rol == "guest" {

		// se conecta al host
		conn, err := net.Dial("tcp", "localhost:"+puerto)
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
