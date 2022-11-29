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

func (a *app) getBlog(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id":           id,
		"title":        "title" + string(id),
		"release_date": "11-26",
		"content":      "nothing",
	})
}
