package db_doc_gen

import "io/ioutil"
import (
	"encoding/json"
	"log"
)

type DbInfo struct {
	DbType   string `json:"db_type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
	IpPort   string `json:"ip_port"`
}

type Config struct {
	Dbinfo       DbInfo   `json:"db_info"`
	Includes     []string `json:"includes"`
	Excludes     []string `json:"excludes"`
	TemplatePath string   `json:"template_path"`
	OutPath      string   `json:"out_path"`
}

func ParseConfigFile(filename string) Config {
	result, err := ioutil.ReadFile(filename)
	if (err != nil) {
		panic("Failed to read the configuration file!")
	}

	var cfg Config
	if err := json.Unmarshal(result, &cfg); err != nil {
		log.Println(err)
		panic("Failed to parse the configuration file!")
	}
	return cfg
}

func CheckConfig(cfg *Config) (bool, string) {
	if IsBlank(cfg.Dbinfo.Schema) {
		return false, "schema can't be empty"
	}
	if IsBlank(cfg.Dbinfo.IpPort) {
		return false, "ip and port can't be empty"
	}
	if IsBlank(cfg.Dbinfo.Username) {
		return false, "ip and port can't be empty"
	}
	if IsBlank(cfg.Dbinfo.Password) {
		return false, "ip and port can't be empty"
	}
	if (IsBlank(cfg.Dbinfo.DbType)) {
		return false, "database type can't be empty"
	}

	if (IsBlank(cfg.OutPath)) {
		cfg.OutPath = cfg.Dbinfo.Schema + "_doc.md"
	}

	return true, "that's right"
}
