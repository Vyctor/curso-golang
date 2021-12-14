package main

import (
	"errors"
	"fmt"
)

func main() {
	// Int
	// int8, int16, int32, int64
	// Se utilizar int ele infere dinamicamente um dos tipos acima
	// tipo uint - Unsigned int, int sem sinal
	numero := 12780000000
	fmt.Println(numero)

	// alias -> rune is a alias to int32
	// alias -> byte is a alias do int8
	var numero3 byte = 128
	fmt.Println(numero3)

	// Float
	// float32, float64
	var numeroReal1 float32 = 123.45
	var numeroReal2 float64 = 12300000000000.45
	numeroReal3 := 123456.78

	fmt.Println(numeroReal1, numeroReal2, numeroReal3)

	// String
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// Booleano
	var booleano bool = true
	fmt.Println(booleano)

	// Erro
	var error error = errors.New("Internal Error!")
	fmt.Println(error)
}
