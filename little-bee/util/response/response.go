package response

import (
	"message/littlebee"
	"message/serialization"
	myerror "util/error"
	"util/i18n"

	"github.com/golang/protobuf/proto"
)

func ResponseErrorWithoutRawError(errorCode myerror.MessageCodeValueEnum, errorDesc string) *littlebee.Response {
	r := &littlebee.Response{
		IsSuccess: false,
		ErrorCode: int32(errorCode),
		ErrorDesc: errorDesc,
		ErrorRaw:  "",
		Result:    nil,
	}
	return r
}

func ResponseErrorByErrorCodeWithoutRawError(errorCode myerror.MessageCodeValueEnum,
	templateData map[string]interface{}, lang string, acceptLanguage string) *littlebee.Response {
	r := &littlebee.Response{
		IsSuccess: false,
		ErrorCode: int32(errorCode),
		ErrorDesc: i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(errorCode)], templateData, lang, acceptLanguage),
		ErrorRaw:  "",
		Result:    nil,
	}
	return r
}

func ResponseError(errorCode myerror.MessageCodeValueEnum, errorDesc string, err error) *littlebee.Response {
	r := &littlebee.Response{
		IsSuccess: false,
		ErrorCode: int32(errorCode),
		ErrorDesc: errorDesc,
		ErrorRaw:  err.Error(),
		Result:    nil,
	}
	return r
}

func ResponseErrorByErrorCode(errorCode myerror.MessageCodeValueEnum,
	templateData map[string]interface{}, lang string, acceptLanguage string, err error) *littlebee.Response {
	r := &littlebee.Response{
		IsSuccess: false,
		ErrorCode: int32(errorCode),
		ErrorDesc: i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(errorCode)], templateData, lang, acceptLanguage),
		ErrorRaw:  err.Error(),
		Result:    nil,
	}
	return r
}

func ResponseErrorJsonWithoutRawError(errorCode myerror.MessageCodeValueEnum, errorDesc string) string {
	json, _ := serialization.MarshalToJson(ResponseErrorWithoutRawError(errorCode, errorDesc))
	return json
}

func ResponseErrorJsonByErrorCodeWithoutRawError(errorCode myerror.MessageCodeValueEnum,
	templateData map[string]interface{}, lang string, acceptLanguage string) string {
	json, _ := serialization.MarshalToJson(ResponseErrorByErrorCodeWithoutRawError(errorCode, templateData, lang, acceptLanguage))
	return json
}

func ResponseErrorJson(errorCode myerror.MessageCodeValueEnum, errorDesc string, err error) string {
	json, _ := serialization.MarshalToJson(ResponseError(errorCode, errorDesc, err))
	return json
}

func ResponseErrorJsonByErrorCode(errorCode myerror.MessageCodeValueEnum,
	templateData map[string]interface{}, lang string, acceptLanguage string, err error) string {
	json, _ := serialization.MarshalToJson(ResponseErrorByErrorCode(errorCode, templateData, lang, acceptLanguage, err))
	return json
}

func ResponseSuccess(m proto.Message) *littlebee.Response {
	r := &littlebee.Response{
		IsSuccess: true,
		ErrorCode: int32(myerror.SUCCESS),
		ErrorDesc: "",
		ErrorRaw:  "",
		Result:    serialization.StructToAny(m),
	}
	return r
}

func ResponseSuccessJson(m proto.Message) string {
	json, _ := serialization.MarshalToJson(ResponseSuccess(m))
	return json
}
