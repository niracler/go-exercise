package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user/*action", func(context *gin.Context) {
		context.String(200, "Hello World")
	})
	r.GET("/test/", func(context *gin.Context) {
		p := context.DefaultQuery("p", "1")
		ps := context.DefaultQuery("ps", "5")
		context.String(http.StatusOK, p+ps)
	})
	r.POST("/test", func(context *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
			context.Abort()
		}
		p := context.PostForm("p", "1")
	})
	
	r.Run()
}
