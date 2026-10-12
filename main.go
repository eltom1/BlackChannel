package main

// print  -> printea literal
// printf -> printea con formato

// el compilador de go es RETRO y rompe las bolas si importas cosas q no usas
import (
	"time" // para registrar el tiempo de los msj

	"bufio" // sirve para leer y escribir en el buffer
	//        bufio.NewReader o bufio.NewScanner

	"fmt" // sirve para imprimir en consola
	"os"  // sirve para interactuar con el sistema operativo

	"net" //sirve para crear conexioncondicion := truees de red

	"strings" // sirve para limpiar consola y manipular strings
)

//os.Args[0] -> nombre del programa
//os.Args[1] -> parametro para el host o guest
//os.Args[2] -> parametro para el puerto

// el scanln rompe el flujo conversacional
//porque sirve para palabras o nros unicamente

// crearInput(string) -> string
// crearInput: funcion que detecta la entrada hasta el salto de linea
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

// AutoHost: funcion que se encarga de crear el host automaticamente c
func AutoHost(puerto string, nombre_user string) {

	fmt.Println("Designado Como Host")
	listener, err := net.Listen("tcp", ":"+puerto) //se pone a la escucha del puerto
	if err != nil {
		fmt.Println("Error en la escucha del puerto", err) //si hay error lo imprime
		return
	}
	conn, err := listener.Accept() //acepta la conexion
	if err != nil {
		fmt.Println("Error al aceptar la conexion:", err) //si hay error lo imprime
		return
	}
	fmt.Println("emparejado con el guest")

	//go routine para poder leer los msj del guest xd
	// funciona como un carril paralelo para poder hacer dos cosas al mismo tiempo
	go func() {
		for {
			buffer := make([]byte, 1024) //crea un buffer de 1024 bytes
			n, err := conn.Read(buffer)  //lee el buffer
			if err != nil {
				fmt.Println("Error al leer el mensaje:", err) //si hay error lo imprime
				os.Exit(0)                                    //sale del programa
			}

			fmt.Print(string(buffer[:n])) //imprime el mensaje del guest
		}
	}()

	//bucle para enviar mensajes
	for {

		mensaje := crearInput("") //nueva funcion implementada

		if mensaje == "exit" {
			break
		} //salir del buble

		hora := time.Now()                    //para obtener la hora exacta de envio
		hora_de_envio := hora.Format("15:04") //formato de hora

		full_mensaje := "[" + nombre_user + "]: " + mensaje + "\t\t\t\t" + "[" + hora_de_envio + "] "
		conn.Write([]byte(full_mensaje + "\n"))
	}
	conn.Close() //termina la conexion
}

// AutoGuest: funcion q desig na el rol de guest automaticamente
func AutoGuest(puerto string, nombre_user string) {
	fmt.Println("Designado Como Guest")
	// se conecta al host
	conn, err := net.Dial("tcp", "localhost:"+puerto)
	if err != nil {
		fmt.Println("Error al conectar con el host:", err) // en caso de no poder conectarse
		return
	}

	fmt.Println("¡Conexion exitosaa con el host")

	//goroutine para el guest que lee los msj del host
	go func() {
		for {
			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Error al leer el mensaje:", err)
				os.Exit(0)
			}
			fmt.Print(string(buffer[:n])) //imprime el mensaje del host
		}
	}()

	//el for ahora solo se encarga de enviar msj
	for {
		mensaje := crearInput("")

		if mensaje == "exit" { //salir del buble
			break
		}

		hora := time.Now()                    //para obtener la hora exacta de envio
		hora_de_envio := hora.Format("15:04") //formato de hora

		full_mensaje := "[" + nombre_user + "]: " + mensaje + "\t\t\t\t" + "[" + hora_de_envio + "] "
		conn.Write([]byte(full_mensaje + "\n"))
	}
	conn.Close() //termina la conexion
}

// funcion main orquestadora del prog
func main() {

	fmt.Println("==================BLACK CHANNEL ==================")

	//tipo := crearInput("ingrese el tipo: host o guest:  ") ya no es necesario

	puerto := crearInput("ingrese el puerto:  ")

	nombre_user := crearInput("ingrese su nombre de Usuario: ")

	listener, err := net.Listen("tcp", ":"+puerto) //se pone a la escucha del puerto
	if err == nil {
		listener.Close() //designa host si no hay errore
		AutoHost(puerto, nombre_user)
	} else {
		AutoGuest(puerto, nombre_user) //sino se designa guest
	}
}
