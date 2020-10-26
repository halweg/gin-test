package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "net/http"
)

type LoginForm struct {
    User string `form:"user" binding:"required"`
    Password string `form:"password" binding:"required"`
}

func main() {

    r := gin.Default()

    r.POST("/login", func(c *gin.Context) {

        var form LoginForm
        e := c.ShouldBindWith(&form, binding.Form)

        if e == nil {
            if form.User == "halweg" && form.Password == "12345678" {
                c.JSON(200, gin.H{"status": "you are logged in"})
            } else {
                c.JSON(401, gin.H{"status":"unauthorized"})
            }
        } else {
            c.JSON(500, gin.H{"status":e.Error()})
        }
    })


    r.Any("/form_post", func(c *gin.Context) {
        message := c.PostForm("message")
        name := c.DefaultPostForm("name", "0")

        uid := c.Query("uid")

        c.JSON(http.StatusOK, gin.H{
            "status" : "success",
            "message" : message,
            "name" : name,
            "uid" : uid,
        })
    })

    r.Run()

}
