package middlewares

import (
	"go_blog/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Verify(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.SECRET),
	})
}
