package casbin

import (
	dao "dao/database"
	"fmt"
	"service/user"
	"strconv"
	"strings"
	myerror "util/error"
	"util/i18n"
	"util/log"
	yaml "util/yaml/service"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

var e *casbin.Enforcer = nil

func CasbinInit() {
	adapter, err := gormadapter.NewAdapterByDB(dao.DB)
	if err != nil {
		msg := i18n.GetLocalTextByDefaultLocal(myerror.MessageCodeValueEnum_desc_id[int32(myerror.ERROR_CASBIN_ADAPTER)], nil)
		log.Logger.Fatal(msg, zap.String(myerror.ERROR_DESC_STRING, err.Error()))
	}
	var err1 error
	e, err1 = casbin.NewEnforcer("./rbac_model.conf", adapter)
	if err1 != nil {
		msg := i18n.GetLocalTextByDefaultLocal(myerror.MessageCodeValueEnum_desc_id[int32(myerror.ERROR_CASBIN_NEW_ENFORCER)], nil)
		log.Logger.Fatal(msg, zap.String(myerror.ERROR_DESC_STRING, err1.Error()))
	}
	err2 := e.LoadPolicy()
	if err2 != nil {
		msg := i18n.GetLocalTextByDefaultLocal(myerror.MessageCodeValueEnum_desc_id[int32(myerror.ERROR_CASBIN_LOAD_POLICY)], nil)
		log.Logger.Fatal(msg, zap.String(myerror.ERROR_DESC_STRING, err2.Error()))
	}

	e.AddFunction("isSuper", func(arguments ...interface{}) (i interface{}, e error) {
		if len(arguments) == 1 {
			user := arguments[0].(string)
			return IsAdmin(user), nil
		}
		return nil, fmt.Errorf("superMatch error")
	})
}

//eg: vo is "u-1" or "g-1"
func IsAdmin(v0 string) bool {
	s := strings.Split(v0, "-")
	if s[0] == "u" {
		userId, _ := strconv.Atoi(s[1])
		u, e := user.RetrieveUser(int32(userId))
		if e == nil && u.UserName == yaml.Yaml.Service.SuperUserName {
			return true
		}
	}
	return false
}

func LoadPolicy() {
	err := e.LoadPolicy()
	if err != nil {
		log.Logger.Error("load policy error", zap.String(myerror.ERROR_DESC_STRING, err.Error()))
	}
}

func AddPolicy(sub string, obj string, act string) (bool, error) {
	b, err := e.AddPolicy(sub, obj, act)
	LoadPolicy()
	return b, err
}

func AddPolicies(rules [][]string) (bool, error) {
	b, err := e.AddPolicies(rules)
	LoadPolicy()
	return b, err
}

func RemoveFilteredPolicy(fieldIndex int, fieldValues string) (bool, error) {
	b, err := e.RemoveFilteredPolicy(fieldIndex, fieldValues)
	LoadPolicy()
	return b, err
}

//AddGroupingPolicy adds a role inheritance rule to the current policy
//adds user to role
func AddGroupingPolicy(sub string, groupName string) (bool, error) {
	b, err := e.AddGroupingPolicy(sub, groupName)
	LoadPolicy()
	return b, err
}

func RemoveGroupingPolicy(sub string, groupName string) (bool, error) {
	b, err := e.RemoveGroupingPolicy(sub, groupName)
	LoadPolicy()
	return b, err
}

func GetFilteredPolicy(fieldIndex int, fieldValues string) [][]string {
	return e.GetFilteredPolicy(fieldIndex, fieldValues)
}

func GetFilteredGroupingPolicy(fieldIndex int, fieldValues string) [][]string {
	return e.GetFilteredGroupingPolicy(fieldIndex, fieldValues)
}

func Enforce(sub string, obj string, act string) (bool, error) {
	return e.Enforce(sub, obj, act)
}
