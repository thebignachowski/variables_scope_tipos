package main

import "fmt"

var primeraVariable string
var numero int
var numeroDecima float32
var numeroDecimal float64
var isNot bool = true
var isVer bool

func unaFuncion() {
	primeraVariable = "AcDc"
	numeroDecima = 33.45
	numeroDecimal = 33.56545854554
	numero = 33
	fmt.Println(primeraVariable, numero, numeroDecima, numeroDecimal, isNot, isVer)

	{
		var segundaVariable string
		segundaVariable = "Rosendo Mercado"
		fmt.Println(segundaVariable)
	}

}

func otraFuncion() {
	primeraVariable = "Metalika"
	fmt.Print(primeraVariable, "\n")
	segundaVariable := "Gun & Roses"
	fmt.Print(segundaVariable)
}

func main() {
	unaFuncion()
	otraFuncion()
}
