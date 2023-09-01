package logout

import (
	"go_mongo_api/src/entities"
)

type LogoutGateway interface {
	ExtractTokenMetadata(encoded string) (*entities.AccessDetails, error)
	DeleteAuth(uuid string) (int64, error)
}
