package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *app) getAllBlogs(c *gin.Context) {
	print("connected")
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
