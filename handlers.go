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
	var blogs []model.Blog

	cursor, err := a.Blogs.Find(context.TODO(), bson.D{})
	if err != nil {
		cursor.Close(context.TODO())
		return
	}

	if err = cursor.All(context.TODO(), &blogs); err != nil {
		log.Printf("error while getting all blogs: %v\n", err)
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

func (a *app) CreateComment(c *gin.Context) {
	// get blogID and comment
	blogID, err := primitive.ObjectIDFromHex(c.Param("blogID"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var comment model.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insert into comments array of corresponding blog
	filter := bson.D{{"_id", blogID}}
	update := bson.D{{"$push", bson.D{{"comments", comment}}}}
	result, err := a.Blogs.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Println(err)
	}

	log.Printf("createComment: %d modified\n", result.ModifiedCount)

	c.JSON(http.StatusOK, comment)
}

func (a *app) DeleteBlog(c *gin.Context) {
	oID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Printf("DeleteBlog: err while getting url paramter %v\n", err)
	}
	// delete in db
	filter := bson.D{{"_id", oID}}
	result, err := a.Blogs.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("err while deleting: %v\n", err)
	}

	log.Printf("DeleteBlog: %d deleted\n", result.DeletedCount)

	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
}
