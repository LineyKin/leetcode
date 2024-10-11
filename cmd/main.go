package main

import "fmt"

func main() {
	fmt.Println(GCD(56, 42))
	fmt.Println(GCD(0, 40))
	fmt.Println(GCD(18, 6))
	fmt.Println(GCD(10, 6))
	fmt.Println(GCD(10, 3))
}

func GCD(a, b int16) int16 {

	if a < b {
		a, b = b, a
	}

	if b == 0 {
		return a
	}

	c := a % b

	if c == 0 {
		return b
	}

	return GCD(b, c)
}
