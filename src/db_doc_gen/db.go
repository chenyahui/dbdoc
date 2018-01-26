package db_doc_gen

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type DbManager struct {
	db  *sql.DB
	cfg Config
}

type ColumnInfo struct {
	Name        string
	ColumnType  string
	Description string
}

type TableInfo struct {
	Columns []ColumnInfo
	Name    string
}

func (self *DbManager) Close() {
	if (self.db != nil) {
		self.db.Close()
	}
}
func (self *DbManager) Connect(cfg Config) {

	self.cfg = cfg

	var err error

	self.db, err = sql.Open(cfg.Dbinfo.DbType, cfg.ConnectStr())

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
				Columns: self.getColumnInfo(tableName),
				Name:    tableName,
			})
	}

	return result
}

func (self *DbManager) getColumnInfo(table_name string) []ColumnInfo {
	query := fmt.Sprintf("SELECT column_name,column_type, column_comment "+
		"FROM information_schema.columns "+
		"WHERE table_schema = DATABASE() "+
		"AND table_name='%s' "+
		"ORDER BY ordinal_position", table_name)

	rows, err := self.db.Query(query)
	if err != nil {
		panic("Failed to query columns info")
	}

	defer rows.Close()

	var result []ColumnInfo
	for rows.Next() {
		column := ColumnInfo{}

		rows.Scan(&column.Name, &column.ColumnType, &column.Description)

		result = append(result, column)
	}
	return result
}
