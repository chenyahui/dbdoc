package dbdoc

import (
	"fmt"
)

func connectFactory(db DbInfo) string {
	var query string
	switch db.DbType {
	case "mssql", "sqlserver":
		query = fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
			db.Username, db.Password, db.IpPort, db.Schema)
	default:
		query = fmt.Sprintf("%s:%s@tcp(%s)/%s",
			db.Username, db.Password, db.IpPort, db.Schema)
	}
	return query
}

func showTableFactory(dbType string) string {
	var query string
	switch dbType {
	case "mssql", "sqlserver":
		query = "SELECT Distinct TABLE_NAME FROM information_schema.TABLES"
	default:
		query = "show tables"
	}
	return query
}

func columnInfoFactory(dbType string, tableName string) string {
	var query string
	switch dbType {
	case "mssql", "sqlserver":
		{
			query = `
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
WHERE tbl.TABLE_NAME = N'%s'
`
		}
	default:
		{
			query = `SELECT column_name,column_type, column_comment 
		FROM information_schema.columns 
		WHERE table_schema = DATABASE() 
		AND table_name='%s' 
		ORDER BY ordinal_position`
		}
	}
	return fmt.Sprintf(query, tableName)
}
