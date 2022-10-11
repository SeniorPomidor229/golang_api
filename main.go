package main

import(
	"github.com/gofiber/fiber/v2"
	
	"workspace/configs"
)


func main(){
    app := fiber.New()
  
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    })
  
    configs.ConnectDB()

    app.Listen(":6000")
}