package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.LoadHTMLGlob("templates/*")


    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{
            "Name": "World",
        })
    })

    err := r.Run(":5001")
    if err != nil {
        panic(err)
    }
}

