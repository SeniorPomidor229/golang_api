package main

import(
	"github.com/gofiber/fiber/v2"
	
	"workspace/configs"
    "workspace/routes"
)


func main(){
    app := fiber.New()
  
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{"data": "Hello form bobert"})
    })
  
    routes.UserRoute(app)

    configs.ConnectDB()

    app.Listen(":6000")
}