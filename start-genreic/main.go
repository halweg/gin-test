package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/bind/*action", func(c *gin.Context) {
        c.String(http.StatusOK, "泛绑定")
    })

    r.Run()
}
