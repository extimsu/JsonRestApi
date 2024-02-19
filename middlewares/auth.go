package middlewares

import (
	"net/http"

	"github.com/extimsu/JsonRestApi/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized" + err.Error()})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
