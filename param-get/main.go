package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/test", func(c *gin.Context) {
        fName := c.Query("fname")
        lName := c.Query("lname")

        c.String(200,  fName+":"+ lName)

    })

    r.Run()
}
