package main

// print  -> printea literal
// printf -> printea con formato

// el compilador de go es retrolazo q rompe las bolas si importas cosas q no usas
import (
	"bufio" // sirve para leer y escribir en el buffer
	"os"

	//            bufio.NewReader o bufio.NewScanner

	"fmt" // sirve para imprimir en consola

	"net" //sirve para crear conexioncondicion := truees de red

	"strings" // sirve para limpiar consola y manipular strings
	//"os" sirve para interactuar con el sistema operativo
	//"os/signal" // sirve para manejar senales del sistema operativo
	//"syscall"  // sirve para capturar senales del sistema
)

//os.Args[0] -> nombre del programa
//os.Args[1] -> parametro para el host o guest
//os.Args[2] -> parametro para el puerto

// el scanln rompe el flujo conversacional
//porque sirve para palabras o nros unicamente

func crearInput(texto string) string {
	fmt.Print(texto)

	//var para leer creando el buffer y vinculandolo al teclado
	lector := bufio.NewReader(os.Stdin) //os.Stdin -> entrada estandar del sistema operativo

	// la lectura hasta q encuentre un salto de linea
	input, err := lector.ReadString('\n') //lee hasta q encuentre un salto de linea
	if err != nil {
		fmt.Println("Error al leer el input:", err) // if para detectar caso de error en lectura de teclado
		return ""

	}
	return strings.TrimSpace(input)
}

func main() {

	fmt.Println("== BLACK CHANNEL ==")
	//checkeo primera version
	//chekeo si no tiene menos de dos parametros tira error
	/*
	   if len(os.Args) < 2 {
	       fmt.Println("Error: escribe host o guest ")
	       return
	   }
	*/

	tipo := crearInput("ingrese el tipo: host o guest:  ")
	puerto := crearInput("ingrese el puerto:  ")

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
		fmt.Println("conexion exitosa con el guest")

		for {
			mensaje := crearInput("escribi un msj... ") //nueva funcion implementada

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
			crearInput("el guest dice: ")

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
