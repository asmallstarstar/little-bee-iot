package fsustatus

import (
	"message/littlebee"
	"message/serialization"
	"net/http"
	fs "service/fsustatus"
	myerror "util/error"
	response "util/response"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func RequestFsuStatus(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	l := &littlebee.FsuStatusRequest{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), l); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		r := fs.ResponseFsuStatus(userId.(int32), l)
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
	}
}
