package main

import (
	"fmt"
	"github.com/tietang/dbx"
	"github.com/tietang/go-utils"
)

func main() {
	x := dbx.Database{}
	fmt.Println(x)
	utils.GetAllIP()
}
