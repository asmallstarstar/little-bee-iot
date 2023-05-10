package testengine

import (
	"message/littlebee"
	"testing"
	yaml "util/yaml/service"

	"github.com/gavv/httpexpect"
)

var TestEngine *httpexpect.Expect = nil
var Token string = ""
var testBaseUrl = "http://127.0.0.1"

func GetTestEngine(t *testing.T) (*httpexpect.Expect, string) {
	if TestEngine == nil {
		TestEngine = httpexpect.New(t, testBaseUrl+yaml.Yaml.Service.Host)
		r := TestEngine.POST("/little-bee/public/user/login").
			WithJSON(&littlebee.LoginRequest{
				UserName: "admin",
				Password: "admin",
			}).Expect().JSON().Object().Raw()
		temp := (r["result"]).(map[string]interface{})
		Token = (temp["token"]).(string)
	}
	return TestEngine, Token
}
