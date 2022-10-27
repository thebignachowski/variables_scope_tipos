package main

import "fmt"

var nombre string = "Nacho"
var apellido string = `Rodríguez`
var nombre_completo string = nombre + " " + apellido

func main() {
	fmt.Print(nombre_completo)
}
