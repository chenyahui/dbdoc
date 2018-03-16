package dbport

import (
	"github.com/chenyahui/dbdoc/dbdoc/common"
	"fmt"
)

type MySQLPort struct {
}

func (MySQLPort) Connect(db common.DbInfo) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		db.Username, db.Password, db.IpPort, db.Schema)
}
func (MySQLPort) ListTables() string {
	return "show tables";
}
func (MySQLPort) GetColumnInfo(tableName string) string {
	return fmt.Sprintf(`SELECT column_name,column_type, column_comment 
		FROM information_schema.columns 
		WHERE table_schema = DATABASE() 
		AND table_name='%s' 
		ORDER BY ordinal_position`, tableName)
}
