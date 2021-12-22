package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mysql MysqlConfig `yaml:"MysqlConfig"`
}

type MysqlConfig struct {
	User     string `yaml:"username"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

var APConfig Config

func ReadConfig(file string) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		// logger.Error("%v", err)

	}
	err = yaml.Unmarshal(buf, &APConfig)
	if err != nil {
		// logger.Error("%v", err)

	}
}
