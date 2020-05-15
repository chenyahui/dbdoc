package dbport

import (
	"github.com/chenyahui/dbdoc/common"
)

type DBOperator interface {
	Connect(db common.DbInfo) string
	ListTables() string
	GetColumnInfo(tableName string) string
}

var registry = map[string]DBOperator{}

func init() {
	registry["mysql"] = MySQLPort{}
	registry["mssql"] = SQLServerPort{}
	registry["sqlserver"] = SQLServerPort{}
}

func GetOperatorByName(name string) DBOperator {
	return registry[name]
}
