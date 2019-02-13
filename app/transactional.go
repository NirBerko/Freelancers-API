package app

import (
	"fmt"
	"freelancers/errors"
	"freelancers/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"
	"strings"
)

func Transactional(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rs := GetRequestScope(c)
		rs.SetDB(db)
	}
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rs := GetRequestScope(c)

		fmt.Println(util.HashAndSalt([]byte("m1ub8wx3")))
		fmt.Println(util.ComparePasswords("$2a$04$Omcu.rqT7zq5bHHyti8GZ.CzJXS0dFoVuRJmSMmCzzkLAbXkPAGfa", []byte("m1ub8wx3")))

		authentication := c.GetHeader("Authorization")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(strings.Split(authentication, "Bearer ")[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("VERIFICATION_KEY")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			rs.SetUserID(uint64(claims["id"].(float64)))
		} else {
			errorHandler := errors.Unauthorized(err.Error())
			c.AbortWithStatusJSON(errorHandler.StatusCode(), errorHandler.Message)
		}
	}
}
