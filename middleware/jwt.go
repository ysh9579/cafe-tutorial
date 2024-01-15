package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"hello-cafe/internal/apierror"
	"hello-cafe/internal/internaljwt"
	"hello-cafe/internal/valid"
	"hello-cafe/model/response"
)

func TokenAuthMiddleware(c *gin.Context) {
	strToken := c.Request.Header.Get("access-token")

	if strToken == "" {
		c.AbortWithStatusJSON(response.Failure(apierror.ErrInvalidAccessToken))
		return
	}

	token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(internaljwt.GetSecretKey()), nil
	})

	if !valid.IsNil(token) {
		fmt.Printf("token raw : %s\n", token.Raw)
	}

	if err != nil {
		c.AbortWithStatusJSON(response.Failure(apierror.ErrInvalidAccessToken))
		return
	}

	c.Next()
}
