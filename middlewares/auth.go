package middlewares

import (
	"net/http"

	"github.com/abhijitpattar/gin-rest-go/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	// get the jwt token
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "empty Token, unauthorized"})
		return
	}

	user_id, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid Token, unauthorized"})
		return
	}

	context.Set("userId", user_id)
	context.Next()
}
