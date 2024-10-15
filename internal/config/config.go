package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	App app `yaml:"app"`
}

type app struct {
	Addr      string `yaml:"addr"`
	ApiPrefix string `yaml:"api_prefix"`
	Debug     bool   `yaml:"debug"`
}

var C *Config

func init() {
	configFile := "default.yaml"
	r, err := os.ReadFile(fmt.Sprintf("./configs/%s", configFile))
	if err != nil {
		panic(err)
	}
	config := &Config{}
	err = yaml.Unmarshal(r, config)
	if err != nil {
		panic(err)
	}
	C = config
}
