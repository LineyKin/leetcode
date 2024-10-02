package easy

// https://leetcode.com/problems/climbing-stairs/description/

func climbStairs(n int) int {
	return Fibonacci(n + 1)
}

func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	prev := 1
	prePrev := 1
	var fib int

	for i := 3; i <= n; i++ {
		fib = prePrev + prev
		prePrev = prev
		prev = fib
	}

	return fib
}

func FibonacciRecursive(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 || n == 2 {
		return 1
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}
