package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
    "net/http"
    "time"
)

type DateRequest struct {
    StartDate time.Time `form:"start_date" binding:"required,DateRight" time_format:"2006-01-02"`

    EndDate time.Time `form:"end_date" binding:"required,gtfield=StartDate" time_format:"2006-01-02"`

    Tag string "hello"
}

func DateRight(l validator.FieldLevel) bool {

    if date, ok := l.Field().Interface().(time.Time); ok {
        now := time.Now()
        if date.Unix() > now.Unix() {
            return true
        }
    }

    return false
}

func main() {
    r := gin.Default()

    if v, ok := binding.Validator.Engine().(*validator.Validate); !ok {
        panic("dateRight bind fail")
    } else {
        v.RegisterValidation("DateRight",  DateRight)
    }

    r.GET("/testing", func(c *gin.Context) {
        var p DateRequest

        if err := c.ShouldBind(&p); err != nil {
            c.String(http.StatusBadRequest, "err : %s", err.Error())
            c.Abort()
            return
        }

        c.String(http.StatusOK, "200: %v", p)
    })

    r.Run()
}
