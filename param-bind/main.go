package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Person struct {
    Name string `form:"name"`
    Age string `form:"age"`
    Job string `form:"job"`
}

func main() {
    r := gin.Default()
    r.GET("/testing", testing)
    r.Run()
}

func testing(c *gin.Context)  {
    var person Person

    if err := c.ShouldBind(&person); err != nil {
        c.String(http.StatusOK, "%s", err.Error())
        return
    }

    c.String(http.StatusOK, "%v", person)
}