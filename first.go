package main

import "fmt"

func variableZeroValue()  {
	var a int
	var s string
	fmt.Println(a, s)
}

func main()  {
	fmt.Println("Hello World")
	variableZeroValue()
}