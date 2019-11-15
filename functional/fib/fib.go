package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fibonacci() IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type IntGen func() int

func (g IntGen) Read(p []byte) (n int, err error) {
	next := g()

	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
