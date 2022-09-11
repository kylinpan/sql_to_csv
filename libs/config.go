package libs

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassWord string
	MysqlName     string
	Query         string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfig(filepath string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
}
