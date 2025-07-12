package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    // TODO: Подключи свои роуты и middleware
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    r.Run() // по умолчанию :8080
}