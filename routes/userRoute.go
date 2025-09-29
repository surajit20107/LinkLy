package routes

import (
  "github.com/gofiber/fiber/v2"
  "main/controllers"
)

func UserRoute(router *fiber.App) {
  route := router.Group("/api/user")
  route.Get("/", controllers.LoginUser)
}