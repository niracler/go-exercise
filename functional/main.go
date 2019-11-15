package main

import "go-exercise/functional/fib"

func main() {
	f := fib.Fibonacci()

	fib.PrintFileContents(f)
}
