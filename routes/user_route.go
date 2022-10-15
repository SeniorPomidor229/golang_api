package routes

import (
	"workspace/controllers"
	"workspace/middleware" 

    "github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
    app.Post("/Create/User", controllers.CreateUser) 

	app.Get("/Get/User/:userId", controllers.GetAUser)

	app.Put("/Change/User/:userId", controllers.EditAUser)

	app.Delete("/Delete/User/:userId", controllers.DeleteAUser) 

	app.Get("/Get/Users", controllers.GetAllUsers)

	app.Post("/Token", controllers.CreateToken)

	app.Post("/Create/Message/:userId", middleware.JWTProtected(), controllers.CreateMessage)
}
