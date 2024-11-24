package routes

import (
	"github.com/axolotl-go/Bank/db"
	"github.com/axolotl-go/Bank/models"
	"github.com/axolotl-go/Bank/services"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if error := c.BodyParser(&user); error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user data",
		})
	}

	user.CardNumber = services.GenerateCardNumber()

	if error := db.DB.Create(&user).Error; error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)

}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	id := c.Params("id")

	if error := db.DB.Where("id = ?", id).Find(&user).Error; error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	if error := db.DB.Find(&users).Error; error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get users",
		})
	}

	return c.JSON(users)
}
