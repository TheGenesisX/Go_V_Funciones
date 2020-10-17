package main

import "fmt"

func generadorImpar() func() uint {
	i := uint(1)

	return func() uint {
		var impar = i
		i += 2
		return impar
	}
}

func main() {
	impar := generadorImpar()
	fmt.Println(impar())
	fmt.Println(impar())
	fmt.Println(impar())
	fmt.Println(impar())
	fmt.Println(impar())
}
