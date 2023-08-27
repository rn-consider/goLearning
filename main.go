package main

import (
   "github.com/gin-gonic/gin"
   "github.com/thinkerou/favicon"
)
 

func main() {
	ginServer := gin.Default()
    ginServer.Use(favicon.New("./判决.png"))
    ginServer.GET("hello",func(context *gin.Context){
		context.JSON(200,gin.H{"msg":"hello world!"})
	})
    ginServer.GET("hello",func(context *gin.Context){
		context.JSON(200,gin.H{"msg":"hello world!"})
	}) 
	ginServer.GET("hello",func(context *gin.Context){
		context.JSON(200,gin.H{"msg":"hello world!"})
	})
	ginServer.GET("hello",func(context *gin.Context){
		context.JSON(200,gin.H{"msg":"hello world!"})
	})
	ginServer.Run(":8082")
}