package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int

	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5

	fmt.Println(array1)

	// Array dinâmico de tamanhos
	array2 := [...]int{1, 2, 3, 4, 5, 6, 6, 7, 8, 89}

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}

	fmt.Println(array2)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	// Arrays Internos
	slice3 := make([]float32, 10, 15)
	fmt.Println("- - - - - - - - - - - -")
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 5)

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
