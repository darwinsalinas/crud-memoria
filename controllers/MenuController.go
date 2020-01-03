package controllers

import (
	"fmt"
	"strings"
)

// PrintMenu imprime el menu en pantalla
func PrintMenu() {
	fmt.Println(strings.ToUpper("1- Lista de usuarios"))
	fmt.Println(strings.ToUpper("2- Ver un usuario"))
	fmt.Println(strings.ToUpper("3- Crear usuario"))
	fmt.Println(strings.ToUpper("4- Actualizar usuario"))
	fmt.Println(strings.ToUpper("5- Salir del programa"))
	doTask()
}

func doTask() {
	var opcion string
	fmt.Scanln(&opcion)

	switch opcion {
	case "1":
		listUsers()
		PrintMenu()
	case "2":
		getUserByIndex()
		PrintMenu()

	case "3":
		createUser()
		PrintMenu()
	case "4":
		updateUser()
		PrintMenu()
	case "5":
		break
	default:
		fmt.Println("Selecciona una opcion de la lista")
		PrintMenu()
	}
}

// PrintHeader imprime la cabecera del programa
func PrintHeader() {
	decorador := "*"
	titulo := "* crud de usuario en memoria *"
	fmt.Println(strings.Repeat(decorador, len(titulo)))
	fmt.Println(strings.ToUpper(titulo))
	fmt.Println(strings.Repeat(decorador, len(titulo)))
}
