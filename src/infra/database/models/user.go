package models

type User struct {
	Base     `bson:",inline"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
