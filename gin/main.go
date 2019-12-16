package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message":"PONG",
		})
	})
	r.StaticFS("/static/", http.Dir("static"))
	r.GET("/hh/:name/:id", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name":context.Param("name"),
			"id":context.Param("id"),
		})
	})
	_ = r.Run()
}
