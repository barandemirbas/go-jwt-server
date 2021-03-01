package middlewares

import (
	"github.com/barandemirbas/go-jwt-server/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func IsAuthorized() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
		SigningKey: []byte(config.GetEnv("SECRET_KEY")),
	})
}
