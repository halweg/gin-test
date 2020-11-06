package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {

    r := gin.Default()

    r.Static("/assets", "./assets")
    r.StaticFS("/static", http.Dir("static"))
    r.StaticFile("/favicon.ico", "./favicon.ico")

    r.GET("/ping", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{"message": "hello world"})
    })
    r.Run()
}