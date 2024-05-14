package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array e slices")

	var array1 [5] string
	array1[0] = "Posição 1 "
	fmt.Println(array1)

	array2 := [5]string{"Posição", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int {10, 11, 12, 13, 41, 15, 16, 17, 18, 19}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	// ARRAY INTERNOS
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(len(slice3)) // LENGTH
	fmt.Println(cap(slice3)) // CAPACIDADE

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}