package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type PersonRequest struct {
    Name string `form:"name" binding:"required" `
    Age int `form:"age" binding:"required,gt=10"`
    Birthday string `form:"birthday" binding:"required"`
}

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
