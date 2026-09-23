package main

// print  -> printea literal
// printf -> printea con formato

// el compilador de go es retrolazo q rompe las bolas si importas cosas q no usas

import (
	"fmt" // sirve para imprimir en consola
	"net" //sirve para crear conexioncondicion := truees de red
	//"os" sirve para interactuar con el sistema operativo
	//"os/signal" // sirve para manejar senales del sistema operativo
	//"syscall"  // sirve para capturar senales del sistema
)

//os.Args[0] -> nombre del programa
//os.Args[1] -> parametro para el host o guest
//os.Args[2] -> parametro para el puerto

func main() {
	//checkeo primera version
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

		fmt.Println("conectado como host")

		listener, err := net.Listen("tcp", ":"+puerto) //escucha el puerto
		conn, err := listener.Accept()                 //acepta la conexion
		if err != nil {
			fmt.Println("Error al aceptar la conexion:", err) //si hay error lo imprime
			return
		}
		fmt.Println("Ya esta conectado")

		for {
			var mensaje string
			fmt.Print("escribi un msj...")
			fmt.Scanln(&mensaje)

			if mensaje == "exit" {
				break
			} //salir del buble
			conn.Write([]byte(mensaje + "\n"))
		}
		conn.Close() //termina la conexion

	} else if rol == "guest" {

		// se conecta al host
		conn, err := net.Dial("tcp", "localhost:"+puerto)
		if err != nil {
			fmt.Println("Error al conectar con el host:", err) // en caso de no poder conectarse
			return
		}

		fmt.Println("¡Conexion exitosaa :0 ")

		for {
			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error al leer el mensaje:", err)
				break
			}
			fmt.Print("el host dice: ", string(buffer[:n])) //imprime el mensaje del host

			var mensaje string
			fmt.Print("escribi un msj...")
			fmt.Scanln(&mensaje)

			if mensaje == "exit" { //salir del buble
				break
			}
			conn.Write([]byte(mensaje + "\n"))
		}
		conn.Close() //termina la conexion
	} else {
		fmt.Println("el rol es host o guest")
	}
}
