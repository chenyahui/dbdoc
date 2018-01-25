package db_doc_gen

import "io/ioutil"
import (
	"encoding/json"
	"fmt"
)

type DbInfo struct {
	DbType   string `json:"db_type"`
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	IpPort   string `json:"ip_port"`
}

type Config struct {
	Dbinfo   DbInfo   `json:"db_info"`
	Includes []string `json:"includes"`
	Excludes []string `json:"excludes"`
	TmplPath string
}

func (self *Config) ConnectStr() string {
	var db = self.Dbinfo

	result := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		db.Username, db.Password, db.IpPort, db.DbName)

	return result
}

func ParseConfigFile(filename string) Config {
	result, err := ioutil.ReadFile(filename)
	if (err != nil) {
		panic(err)
	}

	var cfg Config
	if err := json.Unmarshal(result, &cfg); err != nil {
		panic(err)
	}
	return cfg
}
