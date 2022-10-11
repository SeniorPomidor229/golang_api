package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Phone    string             `json:"phone,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	Subs     bool               `json:"subs,omitempty"`
	Anwers   []Message          `json:"answers,omitempty"`
}
