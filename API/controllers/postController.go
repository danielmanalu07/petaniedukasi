package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"petani_edukasi/database"
	"petani_edukasi/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

const uploadPost = "./filePost"

func init() {
	if _, err := os.Stat(uploadPost); os.IsNotExist(err) {
		os.Mkdir(uploadPost, os.ModePerm)
	}
}

func IndexPost(c *fiber.Ctx) error {
	var post []models.Post

	database.DB.Find(&post)

	if len(post) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Post data not found",
		})
	}

	return c.JSON(post)
}

func ShowPost(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	var posts []models.Post
	database.DB.Where("user_id = ?", userID).Find(&posts)

	if len(posts) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Post data not found for the logged-in user",
		})
	}

	return c.JSON(posts)
}

func CreatePost(c *fiber.Ctx) error {
	image, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "image is required",
		})
	}

	filename := fmt.Sprintf("%d%s", time.Now().Unix(), filepath.Ext(image.Filename))

	if err := c.SaveFile(image, filepath.Join(uploadPost, filename)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	content := c.FormValue("content")

	userID := c.Locals("userID").(string)

	var user models.User
	database.DB.Where("id = ?", userID).First(&user)

	post := models.Post{
		ID:      0,
		Content: content,
		Image:   filename,
		UserID:  user.Id,
	}

	database.DB.Create(&post)

	return c.JSON(post)
}

func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post

	if err := database.DB.Find(&post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Post data not found",
		})
	}

	content := c.FormValue("content")
	userID := c.Locals("userID").(string)

	var user models.User
	database.DB.Where("id = ?", userID).First(&user)

	newImage, newImageErr := c.FormFile("image")
	if newImageErr == nil {
		oldImagePath := filepath.Join(uploadPost, post.Image)
		_ = os.Remove(oldImagePath)

		filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(newImage.Filename))
		if err := c.SaveFile(newImage, filepath.Join(uploadPost, filename)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save new image file",
			})
		}

		post.Image = filename
	}

	if content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Content is required",
		})
	}

	if newImageErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Image is required",
		})
	}

	post.Content = content
	post.UserID = user.Id

	database.DB.Save(&post)

	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.Post
	if err := database.DB.Find(&post, id).Error; err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Post data not found",
		})
	}

	database.DB.Delete(&post)

	return c.JSON(fiber.Map{
		"message": "Successfully deleted the post!",
	})
}
