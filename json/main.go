package main

import (
	"fmt"
	"github.com/json-iterator/go"
)

type user struct {
	Id      int    `json:"id, string"`
	Name    string `json:"username"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"-"`
}

func main() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	u := user{Id: 0, Name: "niracler", Age: 0}
	data, err := json.Marshal(&u)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("序列化后的json数据:", string(data))

	u2 := user{}
	data = []byte(`{"id":1,"username":"gg","age":565}`)

	err = json.Unmarshal(data, &u2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("原来的结构体的数据:%+v \n", u)
	fmt.Printf("反序列化后的json数据:%+v \n", u2)
}
