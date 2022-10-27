package main

import "fmt"

const NUMERO_MINIMO_8BYTES int8 = -128
const NUMERO_MAXIMO_8BYTES int8 = 127
const NUMERO_NIMINO_16BYTES int16 = -32768
const NUMERO_MAXIMO_16BYTES int16 = 32767
const NUMERO_MINIMO_32BYTES int32 = -2147483648
const NUMERO_MAXIMO_32BYTES int32 = 2147483647
const NUMERO_MINIMO_64BYTES int64 = -9223372036854775808
const NUMERO_MAXIMO_64BYTES int64 = 9223372036854775807

var numeroMuyGrande int64 = NUMERO_MAXIMO_64BYTES / 2
var otroNumeroMuyGrande int = int(numeroMuyGrande)

var enteroSinSigno uint = uint(numeroMuyGrande)

var numeroFlotante float32
var numeroFlotanteDoblePrecision float64

var numeroImaginario complex64
var numeroImaginarioMasGrande complex128 = 2 + 5i

// Si declaramos una variable como complex la toma como complex128

var tipoByte byte // es un alias para uint8
var tipoRuna rune // es un alias para int32

func main() {
	fmt.Printf("Numero muy grande - %v => %T\n", numeroMuyGrande, numeroMuyGrande)
	fmt.Printf("Numero muy grande - %v => %T\n", otroNumeroMuyGrande, otroNumeroMuyGrande)
	fmt.Printf("Numero muy grande - %v => %T\n", enteroSinSigno, enteroSinSigno)
}
