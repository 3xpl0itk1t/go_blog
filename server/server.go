package server

import (
	"go_blog/config"
	"go_blog/handlers"
	"go_blog/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartServer() {
	app := fiber.New()
	app.Use(logger.New())

	handlers.ConnectToDB()
	defer handlers.DisconnectFromDB()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome To My Blog",
		})
	})
	jwt := middlewares.Verify(config.SECRET)
	app.Post("/signup", handlers.SignupUser)
	app.Get("/login", handlers.Login)
	app.Post("/write", jwt, handlers.Write)
	app.Get("/posts", jwt , handlers.GetPosts)
	app.Listen(":" + config.PORT)

}
