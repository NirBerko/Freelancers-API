package app

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() gin.HandlerFunc {
	return func(c *gin.Context) {
		ac := newRequestScope(c.Request)

		c.Set("Context", ac)
	}
}

func GetRequestScope(c *gin.Context) RequestScope {
	return c.MustGet("Context").(RequestScope)
}

func newRequestScope(request *http.Request) RequestScope {
	requestID := request.Header.Get("X-Request-Id")
	return &requestScope{
		requestID: requestID,
	}
}

func NewJWT(claims jwt.MapClaims, signingKey string, signingMethod ...jwt.SigningMethod) (string, error) {
	var sm jwt.SigningMethod = jwt.SigningMethodHS256
	if len(signingMethod) > 0 {
		sm = signingMethod[0]
	}
	return jwt.NewWithClaims(sm, claims).SignedString([]byte(signingKey))
}
