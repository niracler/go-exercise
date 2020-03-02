package main

import (
	"bufio"
	"fmt"
	"go-exercise/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 200; i++ {
		_, _ = fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("hello.txt")
}

/*
总结

1. 参数在 defer 语句时调用
2. 确保调用在函数结束时调用
3. defer 列表为先进后出
*/
