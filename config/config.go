package config

import (
	"encoding/json"
	"os"
)

type Environment string
var (
	ENV_DEV     Environment = "dev"
	ENV_QA      Environment = "qa"
	ENV_STAG    Environment = "stag"
	ENV_STAGING Environment = "staging"
	ENV_PROD    Environment = "prod"
)

type RedisCfg struct {
	Addr string `toml:"addr" json:"Addr"`
	Psw  string `toml:"psw" json:"Psw"`
	DBNo int    `toml:"dbno" json:"DBNo"`
}

type MysqlCfg struct {
	Host struct {
		Read  string `toml:"read" json:"read"`
		Write string `toml:"write" json:"write"`
	} `toml:"host" json:"Host"`

	Port    int    `toml:"port" json:"Port"`
	User    string `toml:"user" json:"User"`
	Psw     string `toml:"psw" json:"Psw"`
	DbName  string `toml:"dbname" json:"DbName"`
	LogMode bool   `toml:"Logmode" json:"Logmode"`
}

type Config struct {
	Env Environment `json:"env"`
	RedisDefault RedisCfg `json:"redisDefault"`
	Mysql MysqlCfg `json:"mysqlCfg"`
}

func ParseConfig(filepath string, out interface{}) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(out)
	if err != nil {
		panic(err)
	}
}