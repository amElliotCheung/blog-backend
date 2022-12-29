package main

import (
	"blog/model"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *app) GetAllBlogs(c *gin.Context) {
	var blog model.Blog
	var blogs []model.Blog

	cursor, err := a.Blogs.Find(context.TODO(), bson.D{})
	if err != nil {
		cursor.Close(context.TODO())
		return
	}

	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&blog); err != nil {
			log.Println("decoding blog error")
			return
		}
		blogs = append(blogs, blog)
	}

	log.Println("get all blogs")

	c.JSON(http.StatusOK, blogs)
}

func (a *app) GetBlog(c *gin.Context) {

	blog := model.Blog{}

	oID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Println("illegal parameter id")
		return
	}
	result := a.Blogs.FindOne(context.TODO(),
		bson.D{
			{"_id", oID},
		})
	if err = result.Decode(&blog); err != nil {
		log.Println("decoding error")
		return
	}

	log.Println("get blog")

	c.JSON(http.StatusOK, blog)
}

func (a *app) CreateBlog(c *gin.Context) {

	var blog model.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blog.Id = primitive.NewObjectID()
	result, err := a.Blogs.InsertOne(context.TODO(), blog)

	if err != nil {
		log.Println(err)
	}

	log.Printf("createBlog: %d newly created\n", result.InsertedID)

	c.JSON(http.StatusOK, blog)
}
