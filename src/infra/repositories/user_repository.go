package repositories

import (
	"go_mongo_api/src/entities"
	"go_mongo_api/src/infra/database/models"
	"go_mongo_api/src/infra/mappers"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
	Base       *BaseRepository
	userMapper mappers.UserMapper
}

const collection = "users"

func BuildUserRepository(userMapper *mappers.UserMapper) *UserRepository {
	baseRepo := BuildBaseRepository(collection)

	return &UserRepository{baseRepo, *userMapper}
}

func (r *UserRepository) CreateUser(user entities.User) (*entities.User, error) {
	model := r.userMapper.ToPersistence(user)
	_, err := r.Base.Create(model)

	if err != nil {
		return nil, err
	}

	return r.userMapper.ToDomain(model), nil
}

func (r *UserRepository) UpdateUser(user *entities.User) error {
	// err := r.Db.Save(user).Error
	// return err
	return nil
}

func (r *UserRepository) FindUserByEmail(email string) (*entities.User, error) {
	filter := bson.M{"email": email}

	result := r.Base.FindOne(filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var model models.User

	if err := result.Decode(&model); err != nil {
		return nil, err
	}

	return r.userMapper.ToDomain(model), nil
}

func (r *UserRepository) FindUserByID(id entities.UniqueEntityID) (*entities.User, error) {
	result := r.Base.FindOne(id)
	if result.Err() != nil {
		// Tratar erros, como document not found, aqui se necessário
		return nil, result.Err()
	}

	var user entities.User

	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUsersPaginate(pagination entities.PaginationOptions) ([]entities.User, uint64, error) {
	// var userModels []models.User
	// var users []entities.User
	// var count int64

	// // Executa a consulta para recuperar os produtos paginados e o total de registros correspondentes
	// result := r.Db.Offset(pagination.GetOffset()).Limit(pagination.Limit).Find(&userModels)
	// if result.Error != nil {
	// 	return nil, 0, result.Error
	// }

	// for _, userModel := range userModels {
	// 	users = append(users, *r.UserMapper.ToDomain(userModel))
	// }

	// // Executa uma consulta separada para contar o número total de registros correspondentes
	// countResult := r.Db.Model(&models.User{}).Count(&count)
	// if countResult.Error != nil {
	// 	return nil, 0, countResult.Error
	// }

	// return users, uint64(count), nil
	return nil, 0, nil
}

func (r *UserRepository) StartTransaction() error {
	return r.Base.StartTransaction()
}

func (r *UserRepository) CommitTransaction() error {
	return r.Base.CommitTransaction()
}

func (r *UserRepository) CancelTransaction() error {
	return r.Base.CancelTransaction()
}
