package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {

    r := gin.Default()

    r.GET("/json", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status" : "ok",
            "tag" : " <p> this is a p tag </p> ",
        })
    })

    r.GET("pure-json", func(c *gin.Context) {
        c.PureJSON(http.StatusOK, gin.H{
            "status" : "ok",
            "tag" : "<p> this is a p tag </p>",
        })
    })

    r.Run()
}
