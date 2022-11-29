package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	ReleaseDate time.Time `json:"release_date"`
}

var blogs = []Blog{
	{
		Id:          1,
		Title:       "one",
		Content:     "#title\n- one\n- two\n",
		ReleaseDate: time.Now(),
	},
	{
		Id:          2,
		Title:       "two",
		Content:     "#title\n- one\n- two\n",
		ReleaseDate: time.Now(),
	},
	{
		Id:          3,
		Title:       "three",
		Content:     "#title\n- one\n- two\n",
		ReleaseDate: time.Now(),
	},
}

func (a *app) getAllBlogs(c *gin.Context) {
	print("connected")
	c.JSON(http.StatusOK, gin.H{
		"data": blogs,
	})
}

func (a *app) getBlog(c *gin.Context) {
	pid := c.Param("id")
	id, _ := strconv.Atoi(pid)
	blog := blogs[id]
	c.JSON(http.StatusOK, gin.H{
		"data": blog,
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
