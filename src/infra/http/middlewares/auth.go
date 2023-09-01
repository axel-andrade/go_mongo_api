package middlewares

import (
	"fmt"
	"go_mongo_api/src/infra/bootstrap"
	handlers "go_mongo_api/src/infra/handlers"
	"go_mongo_api/src/infra/repositories"
	ERROR "go_mongo_api/src/shared/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize(dependencies *bootstrap.Dependencies) gin.HandlerFunc {
	tokenManagerHandler := new(handlers.TokenManagerHandler)
	dependencies.Invoke(func(h *handlers.TokenManagerHandler) {
		tokenManagerHandler = h
	})

	sessionRepo := new(repositories.SessionRepository)
	dependencies.Invoke(func(r *repositories.SessionRepository) {
		sessionRepo = r
	})

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			fmt.Println("message: authorization not informed")
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR.UNAUTHORIZED})
			c.Abort()
			return
		}

		encodedToken := authHeader[len("Bearer "):]

		tokenAuth, err := tokenManagerHandler.ExtractTokenMetadata(encodedToken)
		if err != nil {
			fmt.Println("error: error in extract token metadata: ", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR.UNAUTHORIZED})
			c.Abort()
			return
		}

		userId, err := sessionRepo.GetAuth(tokenAuth)
		if err != nil {
			fmt.Println("error in get auth: ", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR.UNAUTHORIZED})
			c.Abort()
			return
		}

		// TODO: verificar se o usuario existe no banco de dados
		c.Set("user-id", userId)

		c.Next()
	}
}
