package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Веб-сервер запущен")
	router := gin.Default()

	router.LoadHTMLGlob("*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "GoOrderFlow"})
	})

	router.Run(":8080")
}
