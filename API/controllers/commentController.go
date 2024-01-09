package controllers

import (
	"petani_edukasi/database"
	"petani_edukasi/models"

	"github.com/gofiber/fiber/v2"
)

func IndexComment(c *fiber.Ctx) error {
	var comment []models.Comment

	database.DB.Find(&comment)

	if len(comment) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Data not found",
		})
	}
	return c.JSON(comment)
}

func CreateComment(c *fiber.Ctx) error {
	postID := c.Params("id")

	text := c.FormValue("text")
	if text == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "text is required",
		})
	}

	var user models.User
	userId := c.Locals("userID").(string)

	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	database.DB.Where("id = ?", userId).First(&user)

	var post models.Post
	if err := database.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "post not found",
		})
	}

	comment := models.Comment{
		Text:   text,
		PostID: post.ID,
		UserID: user.Id,
	}

	database.DB.Create(&comment)

	return c.JSON(comment)
}

func UpdateComment(c *fiber.Ctx) error {
	postID := c.Params("id")

	var comment models.Comment
	if err := database.DB.Find(&comment, postID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No comment found!",
		})
	}

	text := c.FormValue("text")
	if text == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "text is required",
		})
	}

	userID := c.Locals("userID").(string)

	var user models.User
	database.DB.Where("id = ?", userID).First(&user)

	if comment.UserID != user.Id {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not allowed to update this comment",
		})
	}

	var post models.Post
	if err := database.DB.Where("id = ?", postID).First(&post).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "post not found",
		})
	}

	comment.Text = text
	comment.UserID = user.Id
	comment.PostID = post.ID

	database.DB.Save(&comment)

	return c.JSON(fiber.Map{
		"message": "update Success",
	})
}

func DeleteComment(c *fiber.Ctx) error {
	postID := c.Params("id")

	var comment models.Comment
	if err := database.DB.Find(&comment, postID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No comment found!",
		})
	}

	userID := c.Locals("userID").(string)

	var user models.User
	database.DB.Where("id = ?", userID).First(&user)

	if comment.UserID != user.Id {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "You are not allowed to delete this comment",
		})
	}

	database.DB.Delete(&comment)

	return c.JSON(fiber.Map{
		"message": "Deleted Successfully",
	})
}
