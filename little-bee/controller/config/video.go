package config

import (
	"message/littlebee"
	"message/serialization"
	"net/http"
	video "service/config"
	"strconv"
	myerror "util/error"
	response "util/response"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CreateVideo(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	l := &littlebee.Video{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), l); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		r, err := video.CreateVideo(l, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func RetrieveVideo(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if videoId, err := strconv.Atoi(c.Param("video-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		r, err := video.RetrieveVideo(int32(videoId))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func UpdateVideo(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	l := &littlebee.Video{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), l); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := video.UpdateVideo(l, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func DeleteVideo(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if videoId, err := strconv.Atoi(c.Param("video-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := video.DeleteVideo(videoId, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func DeleteVideoWithFlag(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if videoId, err := strconv.Atoi(c.Param("video-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := video.DeleteVideoWithFlag(videoId, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}
