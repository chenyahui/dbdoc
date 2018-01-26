package db_doc_gen

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
)

type DbManager struct {
	db  *sql.DB
	cfg Config
}

type ColumnInfo struct {
	ColumnName  string
	ColumnType  string
	Description string
}

type TableInfo struct {
	Columns   []ColumnInfo
	TableName string
}

func (self *DbManager) Close() {
	if (self.db != nil) {
		self.db.Close()
	}
}
func (self *DbManager) Connect(cfg Config) {

	self.cfg = cfg

	var err error

	self.db, err = sql.Open(cfg.Dbinfo.DbType, connectFactory(self.cfg.Dbinfo))

	if (err != nil) {
		panic("Failed to open database")
	}

	err = self.db.Ping()
	if (err != nil) {
		panic("Failed to connect database")
	}
}

func (self *DbManager) GetTablesInfo() []TableInfo {
	tables := self.filterTables()

	var result []TableInfo

	for _, tableName := range tables {
		result = append(result,
			TableInfo{
				Columns:   self.getColumnInfo(tableName),
				TableName: tableName,
			})
	}

	return result
}

func (self *DbManager) getColumnInfo(tableName string) []ColumnInfo {
	query := columnInfoFactory(self.cfg.Dbinfo.DbType, tableName)
	rows, err := self.db.Query(query)
	if err != nil {
		panic("Failed to query columns info")
	}

	defer rows.Close()

	var result []ColumnInfo
	for rows.Next() {
		column := ColumnInfo{}

		rows.Scan(&column.ColumnName, &column.ColumnType, &column.Description)

		result = append(result, column)
	}
	return result
}
