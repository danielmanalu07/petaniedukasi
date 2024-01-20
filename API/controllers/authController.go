package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"petani_edukasi/database"
	"petani_edukasi/models"
	"time"

	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

const (
	uploadPathKTP  = "./ktpfiles"
	uploadPathFoto = "./fotofiles"
)

func init() {
	if _, err := os.Stat(uploadPathKTP); os.IsNotExist(err) {
		os.Mkdir(uploadPathKTP, os.ModePerm)
	}

	if _, err := os.Stat(uploadPathFoto); os.IsNotExist(err) {
		os.Mkdir(uploadPathFoto, os.ModePerm)
	}
}

func Register(c *fiber.Ctx) error {

	name := c.FormValue("name")
	address := c.FormValue("address")
	age := c.FormValue("age")
	hp := c.FormValue("hp")
	email := c.FormValue("email")
	passwords := c.FormValue("password")

	existingUser := models.User{}
	result := database.DB.Where("email = ?", email).First(&existingUser)
	if result.RowsAffected > 0 {
		c.Status(fiber.StatusConflict)
		return c.JSON(fiber.Map{
			"message": "email already exists",
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(passwords), 14)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "couldn't register",
		})
	}

	ktpFile := c.FormValue("ktp")
	if ktpFile == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to Get KTP File",
		})
	}

	ktpFileName, err := saveFile(c, uploadPathKTP, ktpFile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error saving KTP file",
			"error":   err.Error(),
		})
	}

	fotoFile := c.FormValue("foto")
	if fotoFile == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to Get Foto File",
		})
	}

	fotoFileName, err := saveFile(c, uploadPathFoto, fotoFile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error saving Foto file",
			"error":   err.Error(),
		})
	}

	fmt.Println("KTP File Name:", ktpFileName)
	fmt.Println("Foto File Name:", fotoFileName)

	user := models.User{
		Name:     name,
		Address:  address,
		Age:      age,
		HP:       hp,
		Email:    email,
		Password: password,
		Ktp:      ktpFileName,
		Foto:     fotoFileName,
		Status:   "0",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "error creating user",
			"error":   err.Error(),
		})
	}

	if err := user.ValidateUser(); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "validation error",
			"errors":  err.Error(),
		})
	}

	database.DB.Create(&user)

	user.Password = nil

	return c.JSON(fiber.Map{
		"message": "success",
		"user":    user,
	})
}

func saveFile(c *fiber.Ctx, uploadPath string, fileData string) (string, error) {
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileData))
	filePath := filepath.Join(uploadPath, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return filePath, nil
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	result := database.DB.Where("email = ?", data["email"]).First(&user)

	if result.RowsAffected == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := user.ValidateUser(); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "validation error",
			"errors":  err.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "incorrect email or password",
			})
		}
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "error comparing password",
			"error":   err.Error(),
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "error generating token",
			"error":   err.Error(),
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
