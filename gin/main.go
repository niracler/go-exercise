package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Person struct {
	Name     string `form:"name" binding:"required"`
	Address  string `form:"address"`
	Age      int64 	`form:"age" binding:"required,gt=10"` // 必要字段，并且要大于10
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()

	// 简单 Hello World 请求
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "PONG",
		})
	})

	// 对任意类型请求都做响应
	r.Any("/any", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "你是用什么请求来请求我的呢?",
		})
	})

	// 静态文件夹绑定
	r.StaticFS("/static", http.Dir("static"))

	// 参数作为 url
	r.GET("/hh/:name/:id", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"name": context.Param("name"),
			"id":   context.Param("id"),
		})
	})

	// 泛绑定前缀
	r.GET("/user/*action", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})

	// 获取GET请求的参数,带默认值的参数获取
	r.GET("/test", func(context *gin.Context) {
		p, _ := strconv.ParseInt(context.DefaultQuery("p", "1"), 10, 64)
		ps, _ := strconv.ParseInt(context.DefaultQuery("ps", "5"), 10, 64)
		context.JSON(http.StatusOK, gin.H{
			"p":  p,
			"ps": ps,
		})
	})

	// 获取post请求中，body的信息
	r.POST("/test", func(context *gin.Context) {

		// 拿出来
		bodyBytes, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			context.Abort()
		}
		// 放回去
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		username := context.DefaultPostForm("username", "de_username")
		password := context.DefaultPostForm("password", "de_password")

		context.JSON(http.StatusOK, gin.H{
			"msg":      string(bodyBytes),
			"username": username,
			"password": password,
		})
		//p := context.DefaultPostForm("p", "1")
	})

	// 获取bind参数， 并转换为结构体, 并验证结构体
	r.GET("/to/struct", toStruct)
	r.POST("/to/struct", toStruct)s

	_ = r.Run()
}

func toStruct(context *gin.Context) {
	var person Person
	// 根据请求的content-type来做不同的binding操作
	if err := context.ShouldBind(&person); err == nil {
		context.JSON(http.StatusOK, gin.H{
			"person": fmt.Sprintf("%v", person),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"person": fmt.Sprintf("person bind error: %v", err),
		})
	}
}
