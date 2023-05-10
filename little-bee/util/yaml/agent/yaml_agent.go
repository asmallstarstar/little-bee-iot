package yaml_agent

import (
	"io/ioutil"
	"log"

	yamlV3 "gopkg.in/yaml.v3"
)

const AGENT_YAML_FILE = "agent.yaml"

type YamlAgent struct {
	Agent struct {
		AgentId int32 `yaml:"agent id"`
	} `yaml:"agent"`
	Fsu struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"fsu"`
	Realdata struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"realdata"`
	Log struct {
		AgentFileName    string `yaml:"agent file name"`
		FsuFileName      string `yaml:"fsu file name"`
		RealdataFileName string `yaml:"realdata file name"`
		DriverLoaderName string `yaml:"driver loader file name"`
		MaxSize          int    `yaml:"max size"`
		MaxBackups       int    `yaml:"max backups"`
		MaxAge           int    `yaml:"max age"`
		Level            string `yaml:"level"`
	} `yaml:"log"`

	Database struct {
		Type     string `yaml:"type"`
		Dsn      string `yaml:"dsn"`
		LogLevel int    `yaml:"log level"`
	} `yaml:"database"`

	MQTT struct {
		Broker           string `yaml:"broker"`
		UserName         string `yaml:"user name"`
		Password         string `yaml:"password"`
		RealdataClientId string `yaml:"real data client id"`
	} `yaml:"mqtt"`
}

var Yaml *YamlAgent

func Init() {
	yamlfile, err1 := ioutil.ReadFile(AGENT_YAML_FILE)
	if err1 != nil {
		log.Fatal(err1)
	}
	Yaml = new(YamlAgent)
	err2 := yamlV3.Unmarshal(yamlfile, &Yaml)
	if err2 != nil {
		log.Fatal(err2)
	}
}
