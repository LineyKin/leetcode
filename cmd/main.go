package main

import "fmt"

func main() {
	for i := 0; i < 51; i++ {
		mod3 := i % 3
		if mod3 != 0 {
			fmt.Println(i, mod3)
		}
	}
}
