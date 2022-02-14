package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	FirstName string             `json:"fname,omitempty" validate:"required"`
	LastName  string             `json:"lname,omitempty" validate:"required"`
	UserName  string             `json:"userName,omitempty"`
	Email     string             `json:"email,omitempty" validate:"email"`
	MobileNo  string             `json:"mobileNo,omitempty" validate:"required" binding:"min=7,max=10"`
	Password  string             `json:"password,omitempty"`
	UserType  string             `json:"userType,omitempty"`
	Active    bool               `json:"active,omitempty"`
}
