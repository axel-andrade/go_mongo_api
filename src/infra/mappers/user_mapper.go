package mappers

import (
	"go_mongo_api/src/entities"
	vo "go_mongo_api/src/entities/value_objects"
	"go_mongo_api/src/infra/database/models"
)

type UserMapper struct {
	BaseMapper
}

func BuildUserMapper(baseMapper *BaseMapper) *UserMapper {
	return &UserMapper{BaseMapper: *baseMapper}
}

func (m *UserMapper) ToDomain(model models.User) *entities.User {
	return &entities.User{
		Base:     *m.BaseMapper.toDomain(model.Base),
		Email:    vo.Email{Value: model.Email},
		Name:     vo.Name{Value: model.Name},
		Password: vo.Password{Value: model.Password},
	}
}

func (m *UserMapper) ToPersistence(entity entities.User) models.User {
	return models.User{
		Base:     *m.BaseMapper.toPersistence(entity.Base),
		Email:    entity.Email.Value,
		Name:     entity.Name.Value,
		Password: entity.Password.Value,
	}
}
