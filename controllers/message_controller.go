package controllers

import (
	"context"
    "workspace/configs"
    "workspace/models"
    "workspace/responses"
    "net/http"
    "time"
  
    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)


func CreateMessage(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	var mess models.Message
	defer cancel()

	var user models.User
	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{ Status: http.StatusInternalServerError, Message: "eror", Data: &fiber.Map{"data": err.Error()}})
	}	

	update := bson.M{""}
}