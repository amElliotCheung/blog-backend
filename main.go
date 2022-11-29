package main

import (
	"github.com/gin-gonic/gin"
)

type app struct{}

func (a *app) run() {
	router := gin.Default()
	// https://stackoverflow.com/questions/55347167/gin-contrib-cors-returns-404
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, ResponseType, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	router.GET("/", a.getAllBlogs)
	router.GET("/blogs/:id", a.getBlog)
	router.Run(":8080")
}

func main() {
	a := &app{}
	a.run()
}
