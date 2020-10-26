package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {

    r := gin.Default()

    r.POST("query", func(c *gin.Context) {
        uid := c.Query("uid")
        message := c.DefaultPostForm("message", "no message")
        name := c.PostForm("name")
        list := c.DefaultQuery("list", "null list")

        c.JSON(http.StatusOK, gin.H{
            "uid" : uid,
            "message" : message,
            "name" : name,
            "list" : list,
        })
    })
    r.Run()
}
