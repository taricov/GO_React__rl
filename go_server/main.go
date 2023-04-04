package main

import (
"go_server/routes"
"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
routes.Setup(app)
   app.Listen("localhost:8080")
}
