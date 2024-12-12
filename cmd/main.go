package main

import "fmt"

func main() {
	list := make([]int, 1)

	list = append(list, 2)
	list = append(list, 2)
	list = append(list, 6)
	list = append(list, 2)

	fmt.Println(list)
}
