package main

import (
	"fmt"
	"leetcode/tasks/easy"
	"time"
)

func main() {
	n := 46

	start := time.Now()
	x1 := easy.FibonacciRecursive(n)
	finish := time.Now()

	dur := finish.Sub(start)
	fmt.Println(dur.Nanoseconds())

	start = time.Now()
	x2 := easy.Fibonacci(n)
	finish = time.Now()
	dur = finish.Sub(start)
	fmt.Println(dur.Nanoseconds())

	fmt.Println("x1", x1)
	fmt.Println("x2", x2)
}
