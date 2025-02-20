package middleware

import (
	"os"
	"wisdom/src/utils"

	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware() gin.HandlerFunc {
	var jwtUtils utils.Jwt

	return func(ctx *gin.Context) {
		keyAuth := ctx.GetHeader("Authorization")

		if keyAuth == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
				"reason":  "no token",
			})
			return
		}

		headerToken := keyAuth[len("Bearer "):]

		if headerToken == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
				"reason":  "missing or empty token",
			})
			return
		}

		isValidToken, err := jwtUtils.VerifyToken(headerToken, []byte(os.Getenv("LOGIN_PRIVATE_KEY")))

		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
				"reason":  "token is invalid",
			})
			return
		}

		if !isValidToken.Valid {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "unauthorized",
				"reason":  "token is invalid",
			})
			return
		}

		ctx.Next()
	}
}
