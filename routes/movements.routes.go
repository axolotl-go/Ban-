package routes

import (
	"github.com/axolotl-go/Bank/db"
	"github.com/axolotl-go/Bank/models"
	"github.com/gofiber/fiber/v2"
)

func Deposit(c *fiber.Ctx) error {
	var movement models.Movement
	var user models.User

	id := c.Params("id")

	if error := db.DB.Where("id = ?", id).Find(&user).Error; error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if error := c.BodyParser(&movement); error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	user.Money += movement.Amount

	if err := db.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user's balance",
		})
	}

	return c.JSON(user)

}

func Withdrawal(c *fiber.Ctx) error {
	var movement models.Movement
	var user models.User

	id := c.Params("id")

	if error := db.DB.Where("id =?", id).Find(&user).Error; error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	if error := c.BodyParser(&movement); error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if movement.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Amount must be greater than 0",
		})
	}

	if user.Money < movement.Amount {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Insufficient funds",
		})
	}

	user.Money -= movement.Amount

	if err := db.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user's balance",
		})
	}

	return c.JSON(user)

}
