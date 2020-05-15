package dbport

import (
	"fmt"
	"github.com/chenyahui/dbdoc/common"
)

type SQLServerPort struct {
}

func (SQLServerPort) Connect(db common.DbInfo) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
		db.Username, db.Password, db.IpPort, db.Schema)
}
func (SQLServerPort) ListTables() string {
	return "SELECT Distinct TABLE_NAME FROM information_schema.TABLES"
}
func (SQLServerPort) GetColumnInfo(tableName string) string {
	return fmt.Sprintf(`
				SELECT COLUMN_NAME AS [Output]
				,DATA_TYPE,
					IsNull(Cast(prop.value as varchar(max)),'')
				FROM INFORMATION_SCHEMA.TABLES AS tbl
				INNER JOIN INFORMATION_SCHEMA.COLUMNS AS col ON col.TABLE_NAME = tbl.TABLE_NAME
				INNER JOIN sys.columns AS sc ON sc.object_id = object_id(tbl.table_schema + '.' + tbl.table_name)
				AND sc.NAME = col.COLUMN_NAME
				LEFT JOIN sys.extended_properties prop ON prop.major_id = sc.object_id
				AND prop.minor_id = sc.column_id
				AND prop.NAME = 'MS_Description'
			WHERE tbl.TABLE_NAME = N'%s'`, tableName)
}
