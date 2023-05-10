package middleware

import (
	"errors"
	"fmt"
	"message/littlebee"
	"net/http"
	"regexp"
	"strings"
	"util/casbin"
	myerror "util/error"
	response "util/response"
	yaml_service "util/yaml/service"

	"service/config"

	"github.com/gin-gonic/gin"
)

func Grant() gin.HandlerFunc {
	l, err := config.GetAllActions()
	if err != nil {
		return func(c *gin.Context) {
			lang := c.Request.FormValue("lang")
			acceptLanguage := c.Request.Header.Get("Accept-Language")
			c.Writer.Header().Set("Content-Type", "application/json")
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.FAIL_READ_ACTION_TABLE, nil, lang, acceptLanguage, err))
			c.Abort()
		}
	} else {
		return func(c *gin.Context) {
			url := c.Request.URL.Path
			var a *littlebee.Action = nil
			for _, v := range l.Items {
				if strings.Contains(v.Url, "$") {
					newUrl := "^" + yaml_service.Yaml.Service.UrlRoot + v.Url
					reg := regexp.MustCompile(newUrl)
					m := reg.FindStringIndex(url)
					if m != nil && strings.EqualFold(v.Method, c.Request.Method) {
						a = v
						break
					}
				} else {
					if strings.EqualFold(yaml_service.Yaml.Service.UrlRoot+v.Url, url) && strings.EqualFold(v.Method, c.Request.Method) {
						a = v
						break
					}
				}
			}
			lang := c.Request.FormValue("lang")
			acceptLanguage := c.Request.Header.Get("Accept-Language")
			c.Writer.Header().Set("Content-Type", "application/json")

			if a != nil {
				userId, _ := c.Get("userId")
				sub := fmt.Sprintf("u-%d", userId)
				obj := fmt.Sprintf("d-%d", a.ActionId)
				method := strings.ToUpper(c.Request.Method)
				b, err := casbin.Enforce(sub, obj, method)

				if err != nil {
					c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.FAIL_JUDGMENT_PERMISSION, nil, lang, acceptLanguage, err))
					c.Abort()
				} else {
					if b {
						c.Next()
						if a.ActionType != 4 {
							config.CreateOperationLog(&littlebee.OperationLog{
								UserId:      userId.(int32),
								ActionId:    a.ActionId,
								ActionName:  a.ActionName,
								ActionAlias: a.ActionAlias,
								ActionType:  a.ActionType,
							})
						}
					} else {
						c.String(http.StatusForbidden, "%s", response.ResponseErrorJsonByErrorCode(myerror.FAIL_PERMISSION_DENIED,
							nil, lang, acceptLanguage, errors.New("permission denied")))
						c.Abort()
					}
				}
			} else {
				c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.FAIL_NOTFOUND_ACTION,
					nil, lang, acceptLanguage, errors.New("no matching action found")))
				c.Abort()
			}
		}
	}
}
