package common

import "io/ioutil"
import (
	"encoding/json"
	"log"
	"strings"
	"path/filepath"
)

type Config struct {
	DbInfo       DbInfo  `json:"db_info"`
	Filters      Filters `json:"filters"`
	DocType      string  `json:"doc_type"`
	TemplatePath string  `json:"template_path"`
	OutPath      string  `json:"out_path"`
}

type DbInfo struct {
	DbType   string `json:"db_type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
	IpPort   string `json:"ip_port"`
}

type Filters struct {
	Includes []string `json:"includes"`
	Excludes []string `json:"excludes"`
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
	if IsBlank(cfg.DbInfo.Schema) {
		return false, "schema can't be empty"
	}
	if IsBlank(cfg.DbInfo.IpPort) {
		return false, "ip and port can't be empty"
	}
	if IsBlank(cfg.DbInfo.Username) {
		return false, "ip and port can't be empty"
	}
	if IsBlank(cfg.DbInfo.Password) {
		return false, "ip and port can't be empty"
	}

	if (IsBlank(cfg.DbInfo.DbType)) {
		return false, "database type can't be empty"
	}

	if !InArray([]string{"mysql", "sqlserver", "mssql"}, cfg.DbInfo.DbType) {
		return false, "unsupported database type"
	}

	if IsBlank(cfg.DocType) {
		if !IsBlank(cfg.TemplatePath) &&
			(strings.HasSuffix(cfg.TemplatePath, "doc") ||
				strings.HasSuffix(cfg.TemplatePath, "docx")) {
			cfg.DocType = "word"
		} else {
			cfg.DocType = "plain"
		}
	} else if !InArray([]string{"word", "plain"}, cfg.DocType) {
		return false, "unsupported document type (word/plain)"
	}

	if IsBlank(cfg.OutPath) {
		var ext string
		if !IsBlank(cfg.TemplatePath) {
			ext = filepath.Ext(cfg.TemplatePath)
		} else {
			ext = If(cfg.DocType == "word", "docx", "md").(string)
		}

		cfg.OutPath = cfg.DbInfo.Schema + "_doc." + ext
	}

	return true, "that's right"
}
