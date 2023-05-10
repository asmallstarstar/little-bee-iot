package user

import (
	"errors"
	"fmt"
	"io/ioutil"
	"message/littlebee"
	"message/querycommand"
	"message/serialization"
	"net/http"
	action "service/config"
	"service/user"
	"strings"
	"util/casbin"
	myerror "util/error"
	"util/i18n"
	"util/log"
	response "util/response"

	"strconv"

	yaml "util/yaml/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	u := &littlebee.LoginRequest{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), u); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		r, n, err := user.Login(u)
		if err != nil {
			if n == 0 {
				c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
			} else {
				c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_TOKEN_GENERATE, nil, lang, acceptLanguage, err))
			}
		} else {
			if n != 1 {
				c.String(http.StatusOK, "%s", response.ResponseErrorJsonByErrorCodeWithoutRawError(myerror.ERROR_USERNAME_OR_PASSWORD, nil,
					lang, acceptLanguage))
			} else {
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			}
		}
	}
}

func VerifyUser(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	u := &littlebee.UserVerify{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), u); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		r, _ := user.VerifyUser(u)
		if r == nil {
			c.String(http.StatusOK, "%s", response.ResponseErrorJsonByErrorCodeWithoutRawError(myerror.ERROR_USERNAME_OR_PASSWORD, nil,
				lang, acceptLanguage))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func ModifyPassword(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	u := &littlebee.UserModifyPasword{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), u); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		n, err := user.ModifyPassword(u)
		if n == 1 && err == nil {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseErrorJsonByErrorCode(myerror.FAIL_MODIFY_PASSWORD,
				nil, lang, acceptLanguage, errors.New("modify password failed")))
		}
	}
}

func CreateUser(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	u := &littlebee.User{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), u); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		r, err := user.CreateUser(u, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			insertBaseAction(r.UserId)
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func insertBaseAction(userId int32) {
	a, _ := action.GetActionByName("assembled-menus")
	sub := fmt.Sprintf("u-%d", userId)
	obj := fmt.Sprintf("d-%d", a.ActionId)
	_, err1 := casbin.AddPolicy(sub, obj, a.Method)

	a, _ = action.GetActionByName("user-all-actions")
	sub = fmt.Sprintf("u-%d", userId)
	obj = fmt.Sprintf("d-%d", a.ActionId)
	_, err2 := casbin.AddPolicy(sub, obj, a.Method)

	if err1 != nil {
		log.Logger.Error("inserted base action(assembled-menus) error", zap.String(myerror.ERROR_DESC_STRING, err1.Error()))
	}
	if err2 != nil {
		log.Logger.Error("inserted base action(user-all-actions) error", zap.String(myerror.ERROR_DESC_STRING, err2.Error()))
	}
}

func RetrieveUser(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if userId, err := strconv.Atoi(c.Param("user-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		r, err := user.RetrieveUser(int32(userId))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func UpdateUser(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	u := &littlebee.User{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), u); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		userId, _ := c.Get("userId")
		err := user.UpdateUser(u, userId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func DeleteUser(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if userId, err := strconv.Atoi(c.Param("user-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		deletedByUserId, _ := c.Get("userId")
		err := user.DeleteUser(userId, deletedByUserId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func DeleteUserWithFlag(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if userId, err := strconv.Atoi(c.Param("user-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		deletedByUserId, _ := c.Get("userId")
		err := user.DeleteUserWithFlag(userId, deletedByUserId.(int32))
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_WRITE_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(nil))
		}
	}
}

func QueryUser(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	q := &querycommand.QueryCommand{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), q); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		r, err := user.QueryUser(q)
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
		} else {
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
		}
	}
}

func GetAll(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	r, err := user.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
	} else {
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
	}
}

func GetAllAlias(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	r, err := user.GetAllAlias()
	if err != nil {
		c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
	} else {
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
	}
}

//grant an action to user
func GrantUser(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	p := &littlebee.UserPolicy{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), p); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		sub := fmt.Sprintf("u-%d", p.UserId)
		obj := fmt.Sprintf("d-%d", p.ActionId)
		b, err := casbin.AddPolicy(sub, obj, p.Method)
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_CASBIN_GRANT, nil, lang, acceptLanguage, err))
		} else {
			r := &littlebee.GrantResult{}
			if b {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.SUCCESS_GRANT)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			} else {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.FAIL_GRANT)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			}
		}
	}
}

func BatchGrantUser(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	p := &littlebee.UserPolicyList{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), p); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		var sub = ""
		var userId int32 = 0
		rules := [][]string{}
		for _, v := range p.Items {
			userId = v.UserId
			sub = fmt.Sprintf("u-%d", v.UserId)
			obj := fmt.Sprintf("d-%d", v.ActionId)
			rule := []string{sub, obj, v.Method}
			rules = append(rules, rule)
		}
		addBaseActions(&rules, userId)

		casbin.RemoveFilteredPolicy(0, sub)
		b, err := casbin.AddPolicies(rules)

		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_CASBIN_GRANT, nil, lang, acceptLanguage, err))
		} else {
			r := &littlebee.GrantResult{}
			if b {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.SUCCESS_GRANT)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			} else {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.FAIL_GRANT)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			}
		}
	}
}

func addBaseActions(rules *[][]string, userId int32) {
	a, _ := action.GetActionByName("assembled-menus")
	sub := fmt.Sprintf("u-%d", userId)
	obj := fmt.Sprintf("d-%d", a.ActionId)
	rule := []string{sub, obj, a.Method}
	*rules = append(*rules, rule)

	a, _ = action.GetActionByName("user-all-actions")
	sub = fmt.Sprintf("u-%d", userId)
	obj = fmt.Sprintf("d-%d", a.ActionId)
	rule = []string{sub, obj, a.Method}
	*rules = append(*rules, rule)
}

//grant an action to group
func GrantGroup(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	p := &littlebee.GroupPolicy{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), p); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		sub := fmt.Sprintf("g-%d", p.GroupId)
		obj := fmt.Sprintf("d-%d", p.ActionId)
		b, err := casbin.AddPolicy(sub, obj, p.Method)
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_CASBIN_GRANT, nil, lang, acceptLanguage, err))
		} else {
			r := &littlebee.GrantResult{}
			if b {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.SUCCESS_GRANT)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			} else {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.FAIL_GRANT)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			}
		}
	}
}

func BatchGrantGroup(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	p := &littlebee.GroupPolicyList{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), p); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		var sub = ""
		rules := [][]string{}
		for _, v := range p.Items {
			sub = fmt.Sprintf("g-%d", v.GroupId)
			obj := fmt.Sprintf("d-%d", v.ActionId)
			rule := []string{sub, obj, v.Method}
			rules = append(rules, rule)
		}
		casbin.RemoveFilteredPolicy(0, sub)
		var b bool = true
		var err error = nil
		if len(rules) != 0 {
			b, err = casbin.AddPolicies(rules)
		}

		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_CASBIN_GRANT, nil, lang, acceptLanguage, err))
		} else {
			r := &littlebee.GrantResult{}
			if b {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.SUCCESS_GRANT)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			} else {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.FAIL_GRANT)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			}
		}
	}
}

//get granted actions of group
func GetGroupGranted(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if roleId, err := strconv.Atoi(c.Param("role-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		sub := fmt.Sprintf("g-%d", roleId)
		r := casbin.GetFilteredPolicy(0, sub)
		p := &littlebee.GroupPolicyList{}
		for _, v := range r {
			g := &littlebee.GroupPolicy{}

			s := strings.Split(v[0], "-")
			n, _ := strconv.Atoi(s[1])
			g.GroupId = int32(n)

			s = strings.Split(v[1], "-")
			n, _ = strconv.Atoi(s[1])
			g.ActionId = int32(n)

			g.Method = v[2]
			p.Items = append(p.Items, g)
		}
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(p))
	}
}

//user join group
func JoinGroup(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	g := &littlebee.JoinGroup{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), g); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		sub := fmt.Sprintf("u-%d", g.UserId)
		rule := fmt.Sprintf("g-%d", g.GroupId)
		b, err := casbin.AddGroupingPolicy(sub, rule)
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_CASBIN_GRANT, nil, lang, acceptLanguage, err))
		} else {
			r := &littlebee.GrantResult{}
			if b {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.SUCCESS_JOIN_GROUP)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			} else {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.FAIL_JOIN_GROUP)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			}
		}
	}
}

//remove user from group
func RemoveGroup(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	g := &littlebee.JoinGroup{}
	buf, _ := ioutil.ReadAll(c.Request.Body)

	if err := serialization.UnmarshalFromJson(string(buf), g); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_BODY_PARSING, nil, lang, acceptLanguage, err))
	} else {
		sub := fmt.Sprintf("u-%d", g.UserId)
		rule := fmt.Sprintf("g-%d", g.GroupId)
		b, err := casbin.RemoveGroupingPolicy(sub, rule)
		if err != nil {
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_CASBIN_GRANT, nil, lang, acceptLanguage, err))
		} else {
			r := &littlebee.GrantResult{}
			if b {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.SUCCESS_JOIN_GROUP)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			} else {
				r.IsSuccess = true
				r.Desc = i18n.GetLocalText(myerror.MessageCodeValueEnum_desc_id[int32(myerror.FAIL_JOIN_GROUP)], nil, lang, acceptLanguage)
				c.String(http.StatusOK, "%s", response.ResponseSuccessJson(r))
			}
		}
	}
}

//get all groups for user
func GetUserGroup(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if userId, err := strconv.Atoi(c.Param("user-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		sub := fmt.Sprintf("u-%d", userId)
		r := casbin.GetFilteredGroupingPolicy(0, sub)
		p := &littlebee.JoinGroupList{}
		for _, v := range r {
			g := &littlebee.JoinGroup{}

			s := strings.Split(v[0], "-")
			n, _ := strconv.Atoi(s[1])
			g.UserId = int32(n)

			s = strings.Split(v[1], "-")
			n, _ = strconv.Atoi(s[1])
			g.GroupId = int32(n)

			p.Items = append(p.Items, g)
		}
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(p))
	}
}

//get all actions for user
func GetUserGrant(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if userId, err := strconv.Atoi(c.Param("user-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		sub := fmt.Sprintf("u-%d", userId)
		r := casbin.GetFilteredPolicy(0, sub)
		p := &littlebee.UserPolicyList{}
		for _, v := range r {
			g := &littlebee.UserPolicy{}

			s := strings.Split(v[0], "-")
			n, _ := strconv.Atoi(s[1])
			g.UserId = int32(n)

			s = strings.Split(v[1], "-")
			n, _ = strconv.Atoi(s[1])
			g.ActionId = int32(n)

			g.Method = v[2]

			p.Items = append(p.Items, g)
		}
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(p))
	}
}

//DEACTIVATED
//get all actions for user roles
func GetUserActionsWithRoles(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	if userId, err := strconv.Atoi(c.Param("user-id")); err != nil {
		c.String(http.StatusBadRequest, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_URL_PATH_PARAMETER, nil, lang, acceptLanguage, err))
	} else {
		p := &littlebee.UserPolicyList{}

		sub := fmt.Sprintf("u-%d", userId)
		r := casbin.GetFilteredGroupingPolicy(0, sub)
		for _, v := range r {
			s := strings.Split(v[1], "-")
			n, _ := strconv.Atoi(s[1])
			groupId := int32(n)

			gSub := fmt.Sprintf("g-%d", groupId)
			fr := casbin.GetFilteredPolicy(0, gSub)

			for _, w := range fr {
				g := &littlebee.UserPolicy{}

				ss := strings.Split(w[1], "-")
				nn, _ := strconv.Atoi(ss[1])
				g.ActionId = int32(nn)

				g.Method = w[2]
				g.UserId = (int32)(userId)

				p.Items = append(p.Items, g)
			}

		}
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(p))
	}
}

func AssembledMenus(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	userId, _ := c.Get("userId")
	ms, err1 := action.GetAllMenus()
	mas, err2 := action.GetAllMenuActions()
	allActions, err3 := action.GetAllActions()

	user, _ := user.RetrieveUser(userId.(int32))
	if user.UserName == yaml.Yaml.Service.SuperUserName {
		for _, v := range ms.Items {
			v.IsShow = true
		}
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(ms))
	} else {
		for _, v := range ms.Items {
			v.IsShow = false
		}
		if err1 != nil || err2 != nil || err3 != nil {
			var str string = ""
			if err1 != nil {
				str = str + err1.Error()
			}
			if err2 != nil {
				str = str + err2.Error()
			}
			if err3 != nil {
				str = str + err3.Error()
			}
			c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, errors.New(str)))
		} else {
			actionIdsOfUser := getActionIdsByUserId(userId.(int32))
			for i := 3; i >= 1; i-- {
				is := filter(ms.Items, func(m *littlebee.Menu) bool {
					return m.MenuLevel == (int32)(i)
				})
				for _, v := range is {
					actionNamesOfMenu := make([]string, 0, len(allActions.Items))
					actionIdsOfMenu := make([]int32, 0, len(allActions.Items))
					as := filter(mas.Items, func(ma *littlebee.MenuAction) bool {
						return ma.MenuName == ms.Items[v].MenuName
					})
					for _, k := range as {
						actionNamesOfMenu = append(actionNamesOfMenu, mas.Items[k].ActionName)
					}
					for _, k := range actionNamesOfMenu {
						for _, m := range allActions.Items {
							if m.ActionName == k {
								actionIdsOfMenu = append(actionIdsOfMenu, m.ActionId)
							}
						}
					}

					isShow := isShowMenu(actionIdsOfMenu, actionIdsOfUser)
					children := filter(ms.Items, func(m *littlebee.Menu) bool {
						return m.ParentMenuId == ms.Items[v].MenuId && m.IsShow
					})
					if isShow || len(children) > 0 {
						ms.Items[v].IsShow = true
					} else {
						ms.Items[v].IsShow = false
					}
				}
			}
			c.String(http.StatusOK, "%s", response.ResponseSuccessJson(ms))
		}
	}
}

func isShowMenu(actionIdsOfMenu []int32, actionIdsOfUser []int32) bool {
	for _, r := range actionIdsOfMenu {
		for _, d := range actionIdsOfUser {
			if r == d {
				return true
			}
		}
	}
	return false
}

func getActionIdsByUserId(userId int32) []int32 {
	res := make([]int32, 0)

	//get all actions for user roles
	sub := fmt.Sprintf("u-%d", userId)
	r := casbin.GetFilteredGroupingPolicy(0, sub)
	for _, v := range r {
		s := strings.Split(v[1], "-")
		n, _ := strconv.Atoi(s[1])
		groupId := int32(n)

		gSub := fmt.Sprintf("g-%d", groupId)
		fr := casbin.GetFilteredPolicy(0, gSub)

		for _, w := range fr {
			ss := strings.Split(w[1], "-")
			nn, _ := strconv.Atoi(ss[1])
			res = append(res, (int32)(nn))
		}
	}

	//get all actions for user
	sub = fmt.Sprintf("u-%d", userId)
	r = casbin.GetFilteredPolicy(0, sub)
	for _, v := range r {
		s := strings.Split(v[1], "-")
		n, _ := strconv.Atoi(s[1])
		res = append(res, (int32)(n))
	}

	return res
}

func filter[T any](data []T, f func(T) bool) []int {
	i := make([]int, 0, len(data))
	for k, e := range data {
		if f(e) {
			i = append(i, k)
		}
	}
	return i
}

func GetUserActions(c *gin.Context) {
	lang := c.Request.FormValue("lang")
	acceptLanguage := c.Request.Header.Get("Accept-Language")
	c.Writer.Header().Set("Content-Type", "application/json")

	userId, _ := c.Get("userId")
	userAction := &littlebee.UserActions{}

	a, err := action.GetAllActions()

	u, _ := user.RetrieveUser(userId.(int32))

	if err != nil {
		c.String(http.StatusInternalServerError, "%s", response.ResponseErrorJsonByErrorCode(myerror.ERROR_READ_DATABASE, nil, lang, acceptLanguage, err))
	} else {
		if u.UserName == yaml.Yaml.Service.SuperUserName {
			for _, v := range a.Items {
				userAction.ActionName = append(userAction.ActionName, v.ActionName)
			}
		} else {
			actionIds := getActionIdsByUserId(u.UserId)
			for _, v := range a.Items {
				for _, k := range actionIds {
					if v.ActionId == k {
						userAction.ActionName = append(userAction.ActionName, v.ActionName)
					}
				}
			}
		}
		c.String(http.StatusOK, "%s", response.ResponseSuccessJson(userAction))
	}

}
