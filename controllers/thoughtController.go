package controllers

import (
	"example.com/blog-app/initializers"
	"example.com/blog-app/models"
	"github.com/gofiber/fiber/v2"
)

// CreateThought creates a new thought by using data received from the request body of the http request
func CreateThought(c *fiber.Ctx) error {

	var body struct {
		Title   string
		Content string
		Owner   uint
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	thought := models.Thought{Title: body.Title, Content: body.Content, Owner: body.Owner}

	result := initializers.DB.Create(&thought)

	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}

	return c.JSON(thought)

}

// GetAllThoughts retrieves all thoughts from the database
func GetAllThoughts(c *fiber.Ctx) error {
	thoughts := []models.Thought{}
	initializers.DB.Find(&thoughts)
	return c.JSON(thoughts)
}

func GetThoughtByID(c *fiber.Ctx) error {
	id := c.Params("id")
	thought := models.Thought{}
	result := initializers.DB.First(&thought, id)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}
	return c.JSON(thought)
}

// EditThought Updates a thought by its id or returns an error if the thought does not exist
func EditThought(c *fiber.Ctx) error {
	id := c.Params("id")
	thought := models.Thought{}
	result := initializers.DB.First(&thought, id)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}

	var body struct {
		Title   string
		Content string
		Owner   uint
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	thought.Title = body.Title
	thought.Content = body.Content
	thought.Owner = body.Owner

	result = initializers.DB.Save(&thought)

	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}

	return c.JSON(thought)
}

func DeleteThought(c *fiber.Ctx) error {
	id := c.Params("id")
	thought := models.Thought{}
	result := initializers.DB.First(&thought, id)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}

	initializers.DB.Delete(&thought)

	return c.SendString("Thought successfully deleted")
}
