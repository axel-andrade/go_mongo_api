package login

import (
	"go_mongo_api/src/entities"
)

type LoginGateway interface {
	CreateAuth(userid entities.UniqueEntityID, td *entities.TokenDetails) error
	CompareHashAndPassword(hash string, p string) error
	FindUserByEmail(email string) (*entities.User, error)
	GenerateToken(userid entities.UniqueEntityID) (*entities.TokenDetails, error)
}

type LoginInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutputDTO struct {
	User         entities.User `json:"user"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
}
