package repositories

import (
	"context"
	"go_mongo_api/src/entities"
	database "go_mongo_api/src/infra/database"

	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseRepository struct {
	collection *mongo.Collection
}

func BuildBaseRepository(collectionName string) *BaseRepository {
	db := database.GetDB()
	collection := db.Collection(collectionName)

	return &BaseRepository{collection: collection}
}

func (r *BaseRepository) StartTransaction() error {
	// Not implemented. Maybe use uow pattern for transactions in mongodb
	return nil
}

func (r *BaseRepository) CommitTransaction() error {
	// Not implemented. Maybe use uow pattern for transactions in mongodb
	return nil
}

func (r *BaseRepository) CancelTransaction() error {
	// Not implemented. Maybe use uow pattern for transactions in mongodb
	return nil
}

func (r *BaseRepository) NextEntityID() entities.UniqueEntityID {
	return uuid.NewV4().String()
}

// func (r *BaseRepository) ToUpdate(entity map[string]interface{}) map[string]interface{} {
// 	return r.mapper.ToUpdate(entity)
// }

func (r *BaseRepository) Create(data any) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.TODO(), data)
}

func (r *BaseRepository) FindById(id string) *mongo.SingleResult {
	ctx := context.TODO()
	objectID, _ := primitive.ObjectIDFromHex(id) // Substitua pelo ID desejado
	filter := bson.M{"_id": objectID}
	result := r.collection.FindOne(ctx, filter)

	return result
}

func (r *BaseRepository) FindOne(filter any) *mongo.SingleResult {
	ctx := context.TODO()
	result := r.collection.FindOne(ctx, filter)

	return result
}

func (r *BaseRepository) Count(filter any) (int64, error) {
	return r.collection.CountDocuments(context.TODO(), filter)
}
