package middleware

import (
	"errors"
	"net/http"
	"project_bengkel/auth"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Token() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": http.StatusText(http.StatusUnauthorized)})
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, errors.New(http.StatusText(http.StatusForbidden))
			}

			return []byte(auth.Secret), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Error": http.StatusText(http.StatusForbidden)})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Error": http.StatusText(http.StatusForbidden)})
			return
		}

		exp, _ := claims["exp"].(float64)

		if time.Now().After(time.Unix(int64(exp), 0)) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Error": http.StatusText(http.StatusForbidden)})
			return
		}

		uuid, _ := claims["uid"].(float64)
		ctx.Set("id", uuid)
		ctx.Next()
	}
}
