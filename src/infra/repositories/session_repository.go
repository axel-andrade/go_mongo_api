package repositories

import (
	"context"
	"go_mongo_api/src/entities"
	"go_mongo_api/src/infra/database"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type SessionRepository struct {
	Redis *redis.Client
}

/**
** FetchAuth()aceita o AccessDetails da ExtractTokenMetadatafunção e procura-o no redis. Se o
** registro não for encontrado, pode significar que o token expirou, portanto, um erro é gerado.
**/
func BuildSessionRepository() *SessionRepository {
	return &SessionRepository{Redis: database.GetRedisDB()}
}

func (s *SessionRepository) GetAuth(auth *entities.AccessDetails) (entities.UniqueEntityID, error) {
	userid, err := s.Redis.Get(ctx, auth.AccessUUID).Result()
	if err != nil {
		return "", err
	}

	return userid, nil
}

/**
** Passamos no TokenDetails que contém informações sobre o tempo de expiração dos JWTs e os
** uuids usados ​​na criação dos JWTs. Se o tempo de expiração for atingido para o token de
** atualização ou para o token de acesso , o JWT será excluído automaticamente do Redis.
**/

func (s *SessionRepository) CreateAuth(userid string, td *entities.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := s.Redis.Set(ctx, td.AccessUuid, userid, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := s.Redis.Set(ctx, td.RefreshUuid, userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}

func (s *SessionRepository) DeleteAuth(uuid string) (int64, error) {
	deleted, err := s.Redis.Del(ctx, uuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
