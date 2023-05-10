package config

import (
	"message/littlebee"
	"message/querycommand"
	"message/serialization"
	"net/http"
	signalsi "service/config"
	"strconv"
	myerror "util/error"
	response "util/response"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CreateSignalSi(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	l := &littlebee.SignalSi{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), l); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		r, err := signalsi.CreateSignalSi(l, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func RetrieveSignalSi(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if signalSiId, err := strconv.Atoi(c.Param("signal-si-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		r, err := signalsi.RetrieveSignalSi(int32(signalSiId))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func UpdateSignalSi(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	l := &littlebee.SignalSi{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), l); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := signalsi.UpdateSignalSi(l, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func DeleteSignalSi(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if signalSiId, err := strconv.Atoi(c.Param("signal-si-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := signalsi.DeleteSignalSi(signalSiId, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func QuerySignalSi(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	q := &querycommand.QueryCommand{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), q); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		r, err := signalsi.QuerySignalSi(q)
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func GetAllSignalSis(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	r, err := signalsi.GetAllSignalSis()
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
	} else {
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
	}
}
