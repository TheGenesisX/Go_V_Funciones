package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	n := 10
	i := 0

	for i < n {
		fmt.Println(fibonacci(i))
		i++
	}
}
