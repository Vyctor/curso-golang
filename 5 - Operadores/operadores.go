package main

import "fmt"

func main() {
	// Operadores aritméticos
	fmt.Println("Operadores aritméticos")
	soma := 1 + 1
	subtracao := 1 - 1
	divisao := 100 / 10
	multiplicacao := 10 * 10
	restoDaDivisao := 100 % 10

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	var soma2 int16 = numero1 + numero2
	fmt.Println(soma2)

	// Operadores de atribuição
	fmt.Println("- - - - - - - - - - - - -")
	fmt.Println("Operadores de atribuição")

	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	// Operadores relacionais
	fmt.Println("- - - - - - - - - - - - -")
	fmt.Println("Operadores relacionais")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores lógicos
	fmt.Println("- - - - - - - - - - - - -")
	fmt.Println("Operadores lógicos")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)

	// Operadores unários
	fmt.Println("- - - - - - - - - - - - -")
	fmt.Println("Operadores unários")
	numero := 10
	numero++
	numero += 10
	numero -= 25
	fmt.Println(numero)

	// Operadores ternários
	fmt.Println("- - - - - - - - - - - - -")
	fmt.Println("Operadores ternários")

	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menos ou igual a 5"
	}

	fmt.Println(texto)
}
