package controllers

import(
	"context"
    "workspace/models"
    "workspace/responses"
	"workspace/utils"
    "net/http"
    "time"
  
    "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)


func CreateToken(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.UserDto
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
    }

    if validationErr := validate.Struct(&user); validationErr != nil {
        return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
    }

	err := userCollection.FindOne(ctx, bson.M{"phone": user.Phone}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "erroe", Data: &fiber.Map{"data":err.Error()}})
	} else{
		token, err := utils.GenerateNewAccessToken()
		if err == nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "eror", Data: &fiber.Map{"data":err.Error()}})
		} else{
			return c.Status(http.StatusOK).JSON(responses.TokenResponse{Accsess_token: token, Token_type:"bearer"})
		}
	}

}