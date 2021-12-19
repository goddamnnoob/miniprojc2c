package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var attacks []Attack

func Start() {
	r := gin.Default()
	r.SetTrustedProxies([]string{})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Println("Starting server")
	defer r.Run(":8000")
}
