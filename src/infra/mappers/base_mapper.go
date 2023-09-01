package mappers

import (
	"go_mongo_api/src/entities"
	"go_mongo_api/src/infra/database/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseMapper struct{}

func BuildBaseMapper() *BaseMapper {
	return &BaseMapper{}
}

func (m *BaseMapper) toDomain(model models.Base) *entities.Base {
	return &entities.Base{
		ID:        model.ID.Hex(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (m *BaseMapper) toPersistence(entity entities.Base) *models.Base {

	return &models.Base{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
