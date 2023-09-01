package get_users

import (
	"go_mongo_api/src/entities"
)

type GetUsersGateway interface {
	GetUsersPaginate(pagination entities.PaginationOptions) ([]entities.User, uint64, error)
}

type GetUsersInputDTO struct {
	PaginationOptions entities.PaginationOptions
}

type GetUsersOutputDTO struct {
	Users      []entities.User
	TotalUsers uint64
}
