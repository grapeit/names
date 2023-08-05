package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("*.html")

	r.GET("/", func(c *gin.Context) {
		niceIn := c.Query("nice")
		var nice bool
		if niceIn == "1" {
			nice = true
		} else if niceIn != "0" {
			nice = rand.Intn(2) == 0
		}
		c.HTML(http.StatusOK, "name.html", gin.H{
			"name": generateRandomName(nice)})
	})

	r.GET("/story", func(c *gin.Context) {
		name := c.Query("name")
		c.HTML(http.StatusOK, "story.html", gin.H{
			"name":  name,
			"story": getStory(name),
		})
	})

	r.GET("/runews", func(c *gin.Context) {
		c.HTML(http.StatusOK, "story.html", gin.H{
			"name":  "News",
			"story": getStory("Сренерируй заголовок похожий на новости на русском языке"),
		})
	})

	r.Run(":80")
}
