package get_users

type GetUsersInteractor struct {
	Gateway GetUsersGateway
}

func BuildGetUsersInteractor(g GetUsersGateway) *GetUsersInteractor {
	return &GetUsersInteractor{Gateway: g}
}

func (bs *GetUsersInteractor) Execute(input GetUsersInputDTO) (*GetUsersOutputDTO, error) {
	users, totalUsers, err := bs.Gateway.GetUsersPaginate(input.PaginationOptions)
	if err != nil {
		return nil, err
	}

	return &GetUsersOutputDTO{users, totalUsers}, nil
}
