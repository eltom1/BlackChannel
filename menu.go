package main

import (
	"fmt"
)

func menu() {
	var opccion int

	for {
		fmt.Println("\n== BLACK CHANNEL ==")
		fmt.Println("1. Iniciar como host")
		fmt.Println("2. Iniciar como guest")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")

		//en caso de error
		_, err := fmt.Scan(&opccion)
		if err != nil {
			fmt.Printf("Input invalido")
			continue
		}

		switch opccion {
		case 1:
			fmt.Println("Iniciando como host...")
			main()
		case 2:
			fmt.Println("Iniciando como guest...")
			main()
		case 3:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción inválida. Intente nuevamente.")
		}

	}
}
