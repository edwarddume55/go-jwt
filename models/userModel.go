package models

import{
	"go.mongodb.org/mongo.driver/bson/primitive"
}

type User struct{
	ID          primitive.ObjectID `bso:"_id`
	First_name  *string             `json: "first_name" validate: "required, min=2, max=100" `
	Last_name	*string				`json: "last_name" validate:"validate: "required, min=2, max=100"`
	Password	*string				`json:"pasword" validate: "required, min=8`
	Email		*string				`json: "email" validate: "email, required,`
	Phone		*string				`json: "phone" validate: "required"`
	Token		*string				`json: "token"`
	User_type	*string				`json: "user-type" validate: "required, eq=ADMIN|eq=USER"`
	Refresh_token	*string			`json: "refresh_token"`
	Created_at	time.Time			`json: "created_at"`
	Updated_at	time.Time			`json: "updated_at"`
	User_id		string				`json: "user_id`
}	