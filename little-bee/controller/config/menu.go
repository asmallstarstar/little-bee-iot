package config

import (
	"message/littlebee"
	"message/querycommand"
	"message/serialization"
	"net/http"
	menu "service/config"
	"strconv"
	myerror "util/error"
	response "util/response"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CreateMenu(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	l := &littlebee.Menu{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), l); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		r, err := menu.CreateMenu(l, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func RetrieveMenu(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if menuId, err := strconv.Atoi(c.Param("menu-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		r, err := menu.RetrieveMenu(int32(menuId))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func UpdateMenu(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	l := &littlebee.Menu{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), l); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := menu.UpdateMenu(l, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func DeleteMenu(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if menuId, err := strconv.Atoi(c.Param("menu-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := menu.DeleteMenu(menuId, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func DeleteMenuWithFlag(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if menuId, err := strconv.Atoi(c.Param("menu-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := menu.DeleteMenuWithFlag(menuId, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func QueryMenu(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	q := &querycommand.QueryCommand{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), q); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		r, err := menu.QueryMenu(q)
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func GetAllMenus(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	r, err := menu.GetAllMenus()
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
	} else {
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
	}
}
