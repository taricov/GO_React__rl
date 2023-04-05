package routes

import (
	"go_server/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App){
app.Post("/api/register", controllers.Register)
app.Post("/api/login", controllers.Login)

}
