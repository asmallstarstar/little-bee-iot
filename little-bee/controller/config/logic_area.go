package config

import (
	"message/littlebee"
	"message/querycommand"
	"message/serialization"
	"net/http"
	logicarea "service/config"
	"strconv"
	myerror "util/error"
	response "util/response"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CreateLogicArea(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	l := &littlebee.LogicArea{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), l); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		r, err := logicarea.CreateLogicArea(l, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func RetrieveLogicArea(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if logicAreaId, err := strconv.Atoi(c.Param("logic-area-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		r, err := logicarea.RetrieveLogicArea(int32(logicAreaId))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func UpdateLogicArea(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	l := &littlebee.LogicArea{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), l); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := logicarea.UpdateLogicArea(l, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func DeleteLogicArea(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if logicAreaId, err := strconv.Atoi(c.Param("logic-area-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := logicarea.DeleteLogicArea(logicAreaId, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func QueryLogicArea(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	q := &querycommand.QueryCommand{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), q); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		r, err := logicarea.QueryLogicArea(q)
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func GetAllLogicAreas(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	r, err := logicarea.GetAllLogicAreas()
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
	} else {
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
	}
}
