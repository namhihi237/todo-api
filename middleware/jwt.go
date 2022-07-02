package middleware

import (
	"net/http"
	"strings"
	"todo/pkg/errors"
	"todo/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		var code = errors.SUCCESS
		var err error

		token := getBearerToken(c)
		if token == "" {
			code = errors.UNAUTHORIZED
		} else {
			data, err = util.ParseToken(token)
			if err != nil {
				code = errors.INVALID_TOKEN
			}
		}

		if code != errors.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errors.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Set("user", data)
		c.Next()
	}
}

func getBearerToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		return ""
	}
	return strings.Split(bearerToken, " ")[1]
}
