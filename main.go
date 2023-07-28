package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("*.html")
	r.GET("/name", func(c *gin.Context) {
		niceIn := c.Query("nice")
		var nice bool
		if niceIn == "1" {
			nice = true
		} else if niceIn != "0" {
			nice = rand.Intn(2) == 0
		}
		c.HTML(http.StatusOK, "your-name.html", gin.H{
			"name": generateRandomName(nice)})
	})
	r.Run()
}
