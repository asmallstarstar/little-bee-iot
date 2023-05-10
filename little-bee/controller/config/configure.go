package config

import (
	"message/littlebee"
	"message/querycommand"
	"message/serialization"
	"net/http"
	configure "service/config"
	"strconv"
	myerror "util/error"
	response "util/response"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CreateConfigure(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	l := &littlebee.Configure{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), l); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		r, err := configure.CreateConfigure(l, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func RetrieveConfigure(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if configureId, err := strconv.Atoi(c.Param("configure-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		r, err := configure.RetrieveConfigure(int32(configureId))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func UpdateConfigure(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	l := &littlebee.Configure{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), l); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := configure.UpdateConfigure(l, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func DeleteConfigure(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if configureId, err := strconv.Atoi(c.Param("configure-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := configure.DeleteConfigure(configureId, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func DeleteConfigureWithFlag(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if configureId, err := strconv.Atoi(c.Param("configure-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := configure.DeleteConfigureWithFlag(configureId, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func QueryConfigure(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	q := &querycommand.QueryCommand{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), q); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		r, err := configure.QueryConfigure(q)
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func GetAllConfigures(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	r, err := configure.GetAllConfigures()
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
	} else {
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
	}
}
