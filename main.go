package main

import (
	"github.com/gin-gonic/gin"
)

type app struct {
}

func (a *app) run() {
	router := gin.Default()
	router.GET("/", a.getAllBlogs)
	router.Run(":80")
}

func main() {
	a := &app{}
	a.run()
}
