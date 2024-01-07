package middleware

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

const SecretKey = "secret"

func RequireLogin(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if cookie == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "invalid signature",
			})
		}
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	c.Locals("claims", claims)

	return c.Next()
}
