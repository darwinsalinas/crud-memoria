package controllers

import (
	"crud-memoria/models"
	"fmt"
	"strconv"
	"strings"

	"syreclabs.com/go/faker"
)

var usuarios []models.Usuario

func listUsers() {
	for _, item := range usuarios {
		fmt.Println(item.ID, item.Nombre, item.Apellido)
	}
}

func getUserByIndex() (models.Usuario, int) {
	var index string
	fmt.Scanln(&index)
	i, _ := strconv.Atoi(index)
	if i > len(usuarios) {
		fmt.Println("El usuario no existe")
		return models.Usuario{}, i
	}
	usuario := usuarios[i]

	decorador := "*"
	titulo := "* Usuario del indice " + index + " *"
	fmt.Println(strings.Repeat(decorador, len(titulo)))
	fmt.Println(strings.ToUpper(titulo))
	fmt.Println(strings.Repeat(decorador, len(titulo)))
	fmt.Println(usuario.ID, usuario.Nombre, usuario.Apellido)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	return usuario, i

}

func createUser() {
	var nombre, apellido string
	fmt.Println("Escriba el nombre:")
	fmt.Scanln(&nombre)

	fmt.Println("Escriba el apellido:")
	fmt.Scanln(&apellido)

	usuario := models.Usuario{faker.Number().NumberInt(5), nombre, apellido}
	usuarios = append(usuarios, usuario)
	fmt.Println("usuario creado exitosamente")

}

func updateUser() {
	_, indice := getUserByIndex()
	var nombre, apellido string
	fmt.Println("Escriba el nombre:")
	fmt.Scanln(&nombre)

	fmt.Println("Escriba el apellido:")
	fmt.Scanln(&apellido)

	usuarios[indice].Nombre = nombre
	usuarios[indice].Apellido = apellido

	fmt.Println("Usuario actualizado exitosamente")

}

// FillInitialData rellena la base de datos con datos fake
func FillInitialData(cantidad uint) {
	for i := 0; i <= int(cantidad); i++ {
		usuario := models.Usuario{faker.Number().NumberInt(5), faker.Name().FirstName(), faker.Name().LastName()}
		usuarios = append(usuarios, usuario)
	}
}
