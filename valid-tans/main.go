package main

import (
    "github.com/gin-gonic/gin"
    en2 "github.com/go-playground/locales/en"
    zh2 "github.com/go-playground/locales/zh"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    en_translations "github.com/go-playground/validator/v10/translations/en"
    zh_translations "github.com/go-playground/validator/v10/translations/zh"
    "time"
)

type Person struct {
    Name string `form:"name" validate:"required"`
    Age int `form:"age" validate:"gt=10"`
    JoinDay time.Time
}

func main() {
    valid := validator.New()
    zh := zh2.New()
    en := en2.New()
    uni := ut.New(zh, en)

    r := gin.Default()

    r.GET("/testing", func(c *gin.Context) {
        locale := c.DefaultQuery("locale", "zh")

        tans, _ := uni.GetTranslator(locale)

        switch locale {
        case "zh":
            zh_translations.RegisterDefaultTranslations(valid,tans)
        case "en":
            en_translations.RegisterDefaultTranslations(valid, tans)
        default :
            zh_translations.RegisterDefaultTranslations(valid,tans)
        }

        var person Person

        if err := c.ShouldBind(&person); err != nil {
            c.String(500, err.Error())
            c.Abort()
            return
        }

        if err := valid.Struct(person); err !=nil {
            errs := err.(validator.ValidationErrors)
            errsSlice := []string{}
            for _, e := range errs {
                errsSlice = append(errsSlice, e.Translate(tans))
            }

            c.String(500, "%v",errsSlice)
            c.Abort()
            return
        }

        c.String(200, "%v", person)

    })

    r.Run()
}
