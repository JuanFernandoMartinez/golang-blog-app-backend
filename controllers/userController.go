// this package corresponds to the business logic of the application and is responsible for handling the requests and responses of the user entity
//
// this comment only will be in the overview
package controllers

import (
	"example.com/blog-app/initializers"
	"example.com/blog-app/models"
	"github.com/gofiber/fiber/v2"
)

// CreateUser creates a new user by using data received from the request body of the http request
func CreateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var body struct {
		Username string
		Email    string
		Password string
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	user := models.User{UserName: body.Username, Email: body.Email, Password: body.Password}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}

	return c.JSON(user)
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(c *fiber.Ctx) error {
	users := []models.User{}
	initializers.DB.Find(&users)
	return c.JSON(users)
}

// GetUserByID retrieves a user by its id or returns an error if the user does not exist
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}
	return c.JSON(user)
}

// EditUser Updates a user by its id or returns an error if the user does not exist
func EditUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}

	var body struct {
		Username string
		Email    string
		Password string
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	user.UserName = body.Username
	user.Email = body.Email
	user.Password = body.Password

	initializers.DB.Save(&user)

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		return c.SendString(result.Error.Error())
	}
	initializers.DB.Delete(&user)
	return c.SendString("User successfully deleted")
}
