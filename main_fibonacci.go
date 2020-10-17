package main

import "fmt"

func fibonacci(limite int) {
	a1 := 0
	a2 := 1
	a3 := 1

	for i := 0; i < limite; i++ {
		fmt.Println(a1)
		a1 = a2
		a2 = a3
		a3 = a1 + a2
	}
}

func main() {
	fibonacci(18)
}
