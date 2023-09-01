package signup

import (
	"go_mongo_api/src/entities"
)

type SignupGateway interface {
	CancelTransaction() error
	CreateUser(user entities.User) (*entities.User, error)
	CommitTransaction() error
	EncryptPassword(p string) (string, error)
	FindUserByEmail(email string) (*entities.User, error)
	StartTransaction() error
}

type SignupInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupOutputDTO struct {
	User entities.User
}
