package middleware

import (
	"net/http"

	"util/jwt"

	myerror "util/error"
	response "util/response"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.String(http.StatusUnauthorized, "%s", response.ResponseErrorJsonByErrorCodeWithoutRawError(myerror.ERROR_TOKEN_NOT_FOUND, nil,
				c.Request.FormValue("lang"), c.Request.Header.Get("Accept-Language")))
			c.Abort()
			return
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				c.String(http.StatusUnauthorized, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_TOKEN_EXCEPTION, nil,
					c.Request.FormValue("lang"), c.Request.Header.Get("Accept-Language"), err))
				c.Abort()
				return
			} else {
				c.Set("userId", claims.UserID)
			}
		}
	}
}
