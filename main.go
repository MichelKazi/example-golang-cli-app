package main

import (
	"fmt"
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay"
	"github.com/gin-gonic/gin"
)

func main() {
	logo, _ := cowsay.Say(cowsay.Phrase("Bad Ass MF REST API v4.2.0"), cowsay.Aurora())
	fmt.Println(logo)
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	apiV1 := r.Group("api/v1")
	{
		apiV1.GET("/cow/:phrase", cowHandler)

		apiV1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	apiV2 := r.Group("api/v2")
	{
		apiV2.GET("/cow/:phrase/:title", cowV2Handler)
	}

	r.Run()
}

func cowHandler(c *gin.Context) {
	phrase := c.Param("phrase")
	cow, _ := cowsay.Say(cowsay.Phrase(phrase))
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":  phrase,
		"phrase": cow,
	})
}

func cowV2Handler(c *gin.Context) {
	phrase := c.Param("phrase")
	title := c.Param("title")
	cow, _ := cowsay.Say(cowsay.Phrase(phrase))
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":  title,
		"phrase": cow,
	})
}
