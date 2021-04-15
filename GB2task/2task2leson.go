// Package main provides Fibonacci function for any number.

//Fibonacci(8)

package GB2tsak

import "fmt"

//The main function accept input from user, call the Fibonacci function and print the result
func main() {
	var x int
	fmt.Scan(&x)

	fmt.Println(Fibonacci(x))
}

//The Fibonacci is counting fib. number and return the result
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
