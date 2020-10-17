package main

import "fmt"

func valorMayor(args ...int) int {
	vMayor := args[0]

	for _, v := range args {
		if v > vMayor {
			vMayor = v
		}
	}

	return vMayor
}

func main() {
	a := []int{1, 4, 2, 9, 15, 7, 23, -1}

	fmt.Println(valorMayor(a...))
}
