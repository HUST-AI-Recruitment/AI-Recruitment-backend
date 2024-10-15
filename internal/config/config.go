package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	App      App      `yaml:"app"`
	Jwt      Jwt      `yaml:"jwt"`
	Database Database `yaml:"database"`
}

type App struct {
	Addr      string `yaml:"addr"`
	ApiPrefix string `yaml:"api_prefix"`
	Debug     bool   `yaml:"debug"`
}

type Jwt struct {
	Key    string `yaml:"key"`
	Issuer string `yaml:"issuer"`
}

type Database struct {
	UserName     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	DBName       string `yaml:"db_name"`
	Charset      string `yaml:"charset"`
	ParseTime    bool   `yaml:"parse_time"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

func NewConfig() (*Config, error) {
	configFile := "default.yaml"
	r, err := os.ReadFile(fmt.Sprintf("./configs/%s", configFile))
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(r, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
