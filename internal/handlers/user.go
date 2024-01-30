package handlers

import (
	"File-Processor/config"
	"File-Processor/internal/models"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func FindByCredentials(db *gorm.DB, email, password string) (*models.User, error) {
	var user models.User
	result := db.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, result.Error
	}
	return &user, nil
}

func AuthenticateUser(c *fiber.Ctx, db *gorm.DB, config *config.Config) error {
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := FindByCredentials(db, loginRequest.Email, loginRequest.Password)
	if err != nil {
		if err.Error() == "usuario no encontrado" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Correo electrónico o contraseña no válidos",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.Token.Key))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(models.LoginResponse{
		Token: t,
	})
}
