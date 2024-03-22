package routes

import (
	"learn-middleware/controllers/authController"

	"github.com/gofiber/fiber/v2"
)

func userRouter(app *fiber.App) {

	app.Post("/signup", authController.SignUp)
	app.Post("/login", authController.Login)
	app.Get("/logout", authController.Logout)

}
