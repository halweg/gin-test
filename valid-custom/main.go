package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.GET("/testing", func(c *gin.Context) {
        var p PersonRequest

        if err := c.ShouldBind(&p); err != nil {
            c.String(http.StatusBadRequest, err.Error())
            c.Abort()
            return
        }

        c.String(http.StatusOK, "%v", p)
    })

    r.Run()
}
