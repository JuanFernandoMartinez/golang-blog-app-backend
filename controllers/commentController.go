package controllers

import (
	"example.com/blog-app/initializers"
	"example.com/blog-app/models"
	"github.com/gofiber/fiber/v2"
)

// CreateComment creates a new comment by using data received from the request body of the http request
func CreateComment(c *fiber.Ctx) error {
	var body struct {
		Content   string
		ThoughtID uint
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	thought := models.Comment{Content: body.Content, ThoughtID: body.ThoughtID}

	result := initializers.DB.Create(&thought)

	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}

	return c.JSON(thought)
}

// GetAllComments retrieves all comments from the database
func GetAllComments(c *fiber.Ctx) error {
	comments := []models.Comment{}
	initializers.DB.Find(&comments)
	return c.JSON(comments)
}

func GetCommentByID(c *fiber.Ctx) error {
	id := c.Params("id")
	comment := models.Comment{}
	result := initializers.DB.First(&comment, id)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}
	return c.JSON(comment)
}

func EditComment(c *fiber.Ctx) error {
	id := c.Params("id")
	comment := models.Comment{}
	result := initializers.DB.First(&comment, id)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}

	var body struct {
		Content   string
		ThoughtID uint
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	comment.Content = body.Content
	comment.ThoughtID = body.ThoughtID

	initializers.DB.Save(&comment)

	return c.JSON(comment)
}

func DeleteComment(c *fiber.Ctx) error {
	id := c.Params("id")
	comment := models.Comment{}
	result := initializers.DB.First(&comment, id)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}

	initializers.DB.Delete(&comment)

	return c.SendString("Comment sucessfully deleted")
}
