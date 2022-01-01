package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func calculosMatematicos(a, b int8) (int8, int8) {
	soma := a + b
	subtacao := a - b

	return soma, subtacao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(texto string) {
		fmt.Println(texto)
	}

	f("Êta")

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)

	fmt.Println(resultadoSoma, resultadoSubtracao)
}
