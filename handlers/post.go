package handlers

import (
	"context"
	"go_blog/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Write(c *fiber.Ctx) error {
	var post models.PostModel
	err := c.BodyParser(&post)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	newPost := bson.M{
		"title":   post.Title,
		"content": post.Content,
		"author":  post.Author,
	}

	_, err = blog_collection.InsertOne(context.Background(), newPost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot insert post",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Post Created Successfully",
	})
}

func GetPosts(c *fiber.Ctx) error {
	// Retrieve all posts from the collection
	cursor, err := blog_collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot retrieve posts",
		})
	}
	defer cursor.Close(context.Background())

	var posts []models.PostModel
	if err := cursor.All(context.Background(), &posts); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot decode posts",
		})
	}

	return c.JSON(fiber.Map{
		"posts": posts,
	})
}
