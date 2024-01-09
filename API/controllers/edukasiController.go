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

const uploadPath = "./uploads"

func init() {
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.Mkdir(uploadPath, os.ModePerm)
	}
}

func GetEdukasi(c *fiber.Ctx) error {

	var edukasi []models.Edukasi

	database.DB.Find(&edukasi)

	if len(edukasi) == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Edukasi not found",
		})
	}

	return c.JSON(edukasi)
}

func CreateEdukasi(c *fiber.Ctx) error {
	title := c.FormValue("title")
	description := c.FormValue("description")

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Image is required",
		})
	}

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))

	if err := c.SaveFile(file, filepath.Join(uploadPath, fileName)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	edukasi := models.Edukasi{
		Title:       title,
		Description: description,
		Image:       fileName,
	}
	if err := edukasi.ValidateEdukasi(); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "validation error",
			"errors":  err.Error(),
		})
	}

	database.DB.Create(&edukasi)

	return c.JSON(edukasi)
}

func UpdateEdukasi(c *fiber.Ctx) error {
	id := c.Params("id")

	var edukasi models.Edukasi
	if err := database.DB.Find(&edukasi, id).Error; err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Edukasi not found",
		})
	}

	title := c.FormValue("title")
	description := c.FormValue("description")

	newImage, newImageErr := c.FormFile("image")
	if newImageErr == nil {
		oldImagePath := filepath.Join(uploadPath, edukasi.Image)
		_ = os.Remove(oldImagePath)

		fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(newImage.Filename))
		if err := c.SaveFile(newImage, filepath.Join(uploadPath, fileName)); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save new image file",
			})
		}

		edukasi.Image = fileName
	}

	edukasi.Title = title
	edukasi.Description = description

	if err := edukasi.ValidateEdukasi(); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Validation error",
			"errors":  err.Error(),
		})
	}

	database.DB.Save(&edukasi)

	return c.JSON(edukasi)
}

func DeleteEdukasi(c *fiber.Ctx) error {
	id := c.Params("id")

	var edukasi models.Edukasi
	if err := database.DB.Find(&edukasi, id).Error; err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Edukasi not found",
		})
	}

	database.DB.Delete(&edukasi)

	return c.JSON(fiber.Map{
		"message": "Edukasi deleted successfully",
	})
}
