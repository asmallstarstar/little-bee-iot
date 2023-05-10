package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	yaml "util/yaml/service"

	i18nv2 "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18nv2.Bundle = nil

func InitI18n() {
	bundle = i18nv2.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _, lang := range yaml.Yaml.Service.Local {
		bundle.MustLoadMessageFile(fmt.Sprintf(".\\local\\%v.json", lang))
	}
}

func GetLocalText(id string, templateData map[string]interface{}, lang string, accept string) string {
	localizer := i18nv2.NewLocalizer(bundle, lang, accept)
	return localizer.MustLocalize(&i18nv2.LocalizeConfig{
		DefaultMessage: &i18nv2.Message{
			ID: id,
		},
		TemplateData: templateData,
	})
}

func GetLocalTextByDefaultLocal(id string, templateData map[string]interface{}) string {
	lang, _ := GetLocale()
	localizer := i18nv2.NewLocalizer(bundle, lang, lang)
	return localizer.MustLocalize(&i18nv2.LocalizeConfig{
		DefaultMessage: &i18nv2.Message{
			ID: id,
		},
		TemplateData: templateData,
	})
}

func GetLocale() (string, string) {
	osHost := runtime.GOOS
	defaultLang := "en"
	defaultLoc := "US"
	switch osHost {
	case "windows":
		// Exec powershell Get-Culture on Windows.
		cmd := exec.Command("powershell", "Get-Culture | select -exp Name")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "-")
			lang := langLoc[0]
			loc := langLoc[1]
			return lang, loc
		}
	case "darwin":
		// Exec powershell Get-Culture on Windows.
		cmd := exec.Command("sh", "osascript -e 'user locale of (get system info)'")
		output, err := cmd.Output()
		if err == nil {
			langLocRaw := strings.TrimSpace(string(output))
			langLoc := strings.Split(langLocRaw, "_")
			lang := langLoc[0]
			loc := langLoc[1]
			return lang, loc
		}
	case "linux":
		envlang, ok := os.LookupEnv("LANG")
		if ok {
			langLocRaw := strings.TrimSpace(envlang)
			langLocRaw = strings.Split(envlang, ".")[0]
			langLoc := strings.Split(langLocRaw, "_")
			lang := langLoc[0]
			loc := langLoc[1]
			return lang, loc
		}
	}
	return defaultLang, defaultLoc
}
