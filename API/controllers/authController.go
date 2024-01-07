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
	result := database.DB.Where("email = ?", passwords).First(&existingUser)
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

	ktpFileName := saveFile(c, "ktpfiles", "ktp")
	fotoFileName := saveFile(c, "fotofiles", "foto")

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

	if err := user.ValidateUser(); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "validation error",
			"errors":  err.Error(),
		})
	}

	database.DB.Create(&user)

	user.Password = nil

	return c.JSON(user)
}

func saveFile(c *fiber.Ctx, uploadPath, formFieldName string) string {
	file, err := c.FormFile(formFieldName)
	if err != nil {
		return ""
	}

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
	if err := c.SaveFile(file, filepath.Join(uploadPath, fileName)); err != nil {
		return ""
	}

	return fileName
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
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
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
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
			"message": "couldn't login",
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
