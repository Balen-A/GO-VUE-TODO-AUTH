package controllers

import (
	"Go-Backend/database"
	"Go-Backend/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func TaskCreate(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*claims)
	id := claims.Issuer
	iduint, err := strconv.ParseUint(id, 10, 64)
	task := models.Task{

		Title: data["title"],
		Body:  data["body"],
		User: models.User{
			Id: uint(iduint),
		},
		UserId: uint(iduint),
	}

	if err := database.DB.Create(&task).Error; err != nil {
		c.Status(400)
	}

	return c.JSON(task)
}

func TaskList(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*claims)
	id := claims.Issuer

	var task []models.Task
	database.DB.Where("user_id = ?", id).Find(&task)

	return c.JSON(task)
}
