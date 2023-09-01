package common_ptr

import (
	"go_mongo_api/src/entities"
	"time"
)

type UserFormatted struct {
	ID        entities.UniqueEntityID `json:"id" example:"123" description:"O ID único do usuário"`
	Email     string                  `json:"email" example:"johndoe@example.com" description:"O endereço de email do usuário"`
	Name      string                  `json:"name" example:"John Doe" description:"O nome do usuário"`
	CreatedAt time.Time               `json:"created_at" example:"2022-01-01T00:00:00Z" description:"A data e hora de criação do usuário"`
	UpdatedAt time.Time               `json:"updated_at" example:"2022-01-01T01:00:00Z" description:"A data e hora da última atualização do usuário"`
}

type UserPresenter struct{}

func BuildUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (ptr *UserPresenter) Format(user entities.User) UserFormatted {
	return UserFormatted{
		ID:        user.ID,
		Email:     user.Email.Value,
		Name:      user.Name.Value,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (ptr *UserPresenter) FormatList(users []entities.User) []UserFormatted {
	var usersFormatted []UserFormatted = make([]UserFormatted, 0)

	for _, user := range users {
		usersFormatted = append(usersFormatted, ptr.Format(user))
	}

	return usersFormatted
}
