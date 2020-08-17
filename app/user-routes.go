package app

import (
	"github.com/gofiber/fiber"
	"github.com/juanbelieni/go-simple-api/database"
	"github.com/juanbelieni/go-simple-api/models"
	"github.com/juanbelieni/go-simple-api/utils"
)

func Signin(c *fiber.Ctx) {
	b := struct {
		Name     string `json:"name,notempty"`
		Email    string `json:"email,notempty"`
		Password string `json:"password,notempty"`
	}{}

	if err := c.BodyParser(&b); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body.",
		})
		return
	}

	user := models.User{}

	database.Connection.Where(models.User{
		Email: b.Email,
	}).First(&user)

	if user.ID != 0 {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already registered.",
		})
		return
	}

	hs := utils.GetHashedPassword(b.Password)

	database.Connection.Create(&models.User{
		Name:           b.Name,
		Email:          b.Email,
		HashedPassword: hs,
	})

	c.Status(fiber.StatusCreated).Send()
}

func Login(c *fiber.Ctx) {
	b := struct {
		Email    string `json:"email,notempty"`
		Password string `json:"password,notempty"`
	}{}

	if err := c.BodyParser(&b); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing body.",
		})
		return
	}

	user := models.User{}

	database.Connection.Where(models.User{
		Email:          b.Email,
		HashedPassword: utils.GetHashedPassword(b.Password),
	}).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email or password incorrect.",
		})
		return
	}

	c.JSON(user)
}
