package handlers

import (
	"fmt"
	"go_mongo_api/src/entities"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type TokenManagerHandler struct{}

func BuildTokenManagerHandler() *TokenManagerHandler {
	return &TokenManagerHandler{}
}

func (tmi *TokenManagerHandler) GenerateToken(userid string) (*entities.TokenDetails, error) {
	td := entities.TokenDetails{}

	if err := tmi.configureExpiration(&td); err != nil {
		return nil, err
	}

	if err := tmi.createAccessToken(&td, userid); err != nil {
		return nil, err
	}

	if err := tmi.createRefreshToken(&td, userid); err != nil {
		return nil, err
	}

	return &td, nil
}

func (tmi *TokenManagerHandler) VerifyToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (any, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tmi.getSecretKey()), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (tmi *TokenManagerHandler) ExtractTokenMetadata(encoded string) (*entities.AccessDetails, error) {
	token, err := tmi.VerifyToken(encoded)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		return &entities.AccessDetails{
			AccessUUID: accessUuid,
			UserID:     claims["access_uuid"].(string),
		}, nil
	}

	return nil, err
}

func (tm *TokenManagerHandler) TokenValid(encodedToken string) error {

	token, err := tm.VerifyToken(encodedToken)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return err
	}

	return nil
}

func (tmi *TokenManagerHandler) getSecretKey() string {
	secret := os.Getenv("ACCESS_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (tmi *TokenManagerHandler) createAccessToken(td *entities.TokenDetails, userid string) error {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(tmi.getSecretKey()))
	if err != nil {
		return err
	}

	return nil
}

func (e *TokenManagerHandler) createRefreshToken(td *entities.TokenDetails, userid string) error {
	var err error

	refreshSecret := os.Getenv("REFRESH_SECRET")

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return err
	}

	return nil
}

func (e *TokenManagerHandler) configureExpiration(td *entities.TokenDetails) error {
	expiration, err := strconv.Atoi(os.Getenv("MINUTES_TO_EXPIRATION_TOKEN"))
	if err != nil {
		return fmt.Errorf("error converting expiration to int")
	}

	td.AtExpires = time.Now().Add(time.Minute * time.Duration(expiration)).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	return nil
}
