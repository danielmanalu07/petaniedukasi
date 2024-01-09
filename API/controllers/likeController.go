package controllers

import (
	"petani_edukasi/database"
	"petani_edukasi/models"

	"github.com/gofiber/fiber/v2"
)

func LikePost(c *fiber.Ctx) error {
	userId := c.Locals("userID")
	postId := c.Params("id")

	var existingLike models.Like
	if err := database.DB.Where("user_id = ? AND post_id = ? AND status = 0", userId, postId).First(&existingLike).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "User has already liked the post",
		})
	}

	var user models.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	var post models.Post
	if err := database.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "post not found",
		})
	}

	like := models.Like{
		UserID: user.Id,
		PostID: post.ID,
		Status: 0,
	}

	if err := database.DB.Create(&like).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to like the post",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Liked",
	})
}

func DeleteLikeDislike(c *fiber.Ctx) error {
	id := c.Params("id")
	userId := c.Locals("userID")

	var like models.Like
	if err := database.DB.Where("id = ? AND user_id = ?", id, userId).First(&like).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	database.DB.Delete(&like)

	return c.JSON(fiber.Map{
		"message": "Successfully Deleted",
	})
}

func DislikePost(c *fiber.Ctx) error {
	userId := c.Locals("userID")
	postId := c.Params("id")

	var existingDisLike models.Like
	if err := database.DB.Where("user_id = ? AND post_id = ? AND status = 1", userId, postId).First(&existingDisLike).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "User has already disliked the post",
		})
	}

	var user models.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	var post models.Post
	if err := database.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "post not found",
		})
	}

	like := models.Like{
		UserID: user.Id,
		PostID: post.ID,
		Status: 1,
	}

	if err := database.DB.Create(&like).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to like the post",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Disliked",
	})

}
