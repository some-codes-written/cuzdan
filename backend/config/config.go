package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/Tkanos/gonfig"
	"golang.org/x/exp/slices"
)

const CURRENT_ENV = "dev"

type Config struct {
	Env struct {
		Type    string `json:Type`
		BaseUrl string `json:BaseUrl`
	}

	DatabaseConfig struct {
		Name     string `json:Name`
		Host     string `json:Host`
		Username string `json:Username`
		Password string `json:Password`
		Port     string `json:Port`
	}

	NameValues []NameValue
}

type NameValue struct {
	ID    string `json:"ID"`
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

func GetConfig(params ...string) Config {
	conf := Config{}
	env := CURRENT_ENV
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("./%s_config.json", env)
	if gonfig.GetConf(fileName, &conf) != nil {
		panic("config file not found")
	}
	return conf
}

func GetCommand(valueId string) NameValue {
	if valueId == "" {
		panic("value id cannot be null")
	}
	values := GetConfig().NameValues
	idx := slices.IndexFunc(values, func(value NameValue) bool { return value.ID == valueId })
	if idx == -1 {
		panic("value not found")
	}
	return values[idx]
}

func ConfigFileExists() bool {
	fileName := fmt.Sprintf("../config/%s_config.json", CURRENT_ENV)
	_, err := os.Stat(fileName)

	return !errors.Is(err, os.ErrNotExist)
}
