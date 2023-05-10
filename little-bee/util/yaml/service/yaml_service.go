package yaml_service

import (
	"io/ioutil"
	"log"

	yamlV3 "gopkg.in/yaml.v3"
)

const SERVICE_YAML_FILE = "service.yaml"

type YamlService struct {
	Service struct {
		Host            string   `yaml:"host"`
		UrlRoot         string   `yaml:"url root"`
		WebRunMode      string   `yaml:"web run mode"`
		JwtSecret       string   `yaml:"jwt secret"`
		TokenExpireTime int32    `yaml:"token expire time"`
		Local           []string `yaml:"local"`
		SuperUserName   string   `yaml:"super user name"`
	} `yaml:"service"`
	Log struct {
		FileName   string `yaml:"file name"`
		MaxSize    int    `yaml:"max size"`
		MaxBackups int    `yaml:"max backups"`
		MaxAge     int    `yaml:"max age"`
		Level      string `yaml:"level"`
	} `yaml:"log"`

	Database struct {
		Type     string `yaml:"type"`
		Dsn      string `yaml:"dsn"`
		LogLevel int    `yaml:"log level"`
	} `yaml:"database"`

	MQTT struct {
		Broker   string `yaml:"broker"`
		UserName string `yaml:"user name"`
		Password string `yaml:"password"`
		ClientId string `yaml:"client id"`
	} `yaml:"mqtt"`
}

var Yaml *YamlService

func Init() {
	yamlfile, err1 := ioutil.ReadFile(SERVICE_YAML_FILE)
	if err1 != nil {
		log.Fatal(err1)
	}
	Yaml = new(YamlService)
	err2 := yamlV3.Unmarshal(yamlfile, &Yaml)
	if err2 != nil {
		log.Fatal(err2)
	}
}
