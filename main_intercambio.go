package main

import "fmt"

func intercambio(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	var a int
	var b int

	fmt.Println("Ingrese los valores de a y b: ")
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println("Valor de 'a' antes del intercambio:", a)
	fmt.Println("Valor de 'b' antes del intercambio:", b)

	intercambio(&a, &b)

	fmt.Println("Valor de 'a' despues del intercambio:", a)
	fmt.Println("Valor de 'b' despues del intercambio:", b)
}
