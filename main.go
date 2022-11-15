package main

import (
	"github.com/gin-gonic/gin"
)

type app struct {
}

func (a *app) run() {
	router := gin.Default()
	router.GET("/", a.getAllBlogs)
	router.Run(":8080")
}

func main() {
	a := &app{}
	a.run()
}
