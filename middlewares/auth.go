package middleware_auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := ctx.Request.Cookie("token")
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		fmt.Println(tokenString)
		if tokenString != nil {
			token, err := jwt.Parse(tokenString.Value, func(tok *jwt.Token) (interface{}, error) {
				fmt.Printf("unexpected signing method: %v", tok.Method)
				if tok.Method != jwt.SigningMethodHS256 {
					return nil, fmt.Errorf("unexpected signing method: %v", tok.Method)
				}

				return []byte("secretkey"), nil
			})

			if err != nil {
				log.Println(err)
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx.Next()
				return
			} else {
				fmt.Println(err)

				return
			}
		}

		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": "false",
			"error":  "",
		})
	}
}
