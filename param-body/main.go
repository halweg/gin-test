package main

import (
    "bytes"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "net/http"
)

func main() {
    r := gin.Default()
    r.POST("/test", func(c *gin.Context) {
        body, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
            c.String(http.StatusBadRequest, err.Error())
            c.Abort()
        }
        c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

        fName := c.PostForm("first_name")
        lName := c.PostForm("last_name")

        c.String(http.StatusOK, "%s, %s, %s" , fName,lName,string(body))
    })

    r.Run()
}
