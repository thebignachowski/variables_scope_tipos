package main

import "fmt"

func registrarse() {
	var isRegistrado bool = true
	var isEmailConfirmado bool
	var isAccesoCompleto = isRegistrado && isEmailConfirmado
	var isParcialAcceso = isRegistrado || isEmailConfirmado

	fmt.Printf("Está registrado - %v => %T\n", isRegistrado, isRegistrado)
	fmt.Printf("Está el mail registrado - %v => %T\n", isEmailConfirmado, isEmailConfirmado)

	fmt.Printf("Acceso Completo - %v => %T\n", isAccesoCompleto, isAccesoCompleto)
	fmt.Printf("Acceso Parcial - %v => %T\n", isParcialAcceso, isParcialAcceso)
}

func main() {
	registrarse()
}
