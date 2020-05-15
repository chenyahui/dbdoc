package process

import (
	"database/sql"
	"github.com/chenyahui/dbdoc/common"
	"github.com/chenyahui/dbdoc/dbport"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

type DbManager struct {
	db       *sql.DB
	cfg      common.Config
	operator dbport.DBOperator
}

func (self *DbManager) Close() {
	if self.db != nil {
		self.db.Close()
	}
}

func (self *DbManager) Connect(cfg common.Config) {
	self.cfg = cfg
	self.operator = dbport.GetOperatorByName(cfg.DbInfo.DbType)

	var err error

	self.db, err = sql.Open(cfg.DbInfo.DbType, self.operator.Connect(cfg.DbInfo))

	if err != nil {
		panic("Failed to open database")
	}

	err = self.db.Ping()
	if err != nil {
		panic("Failed to connect database")
	}
}

func (self *DbManager) GetTablesInfo() []common.TableInfo {
	tables := self.filterTables()

	var result []common.TableInfo

	for _, tableName := range tables {
		result = append(result,
			common.TableInfo{
				Columns:   self.getColumnInfo(tableName),
				TableName: tableName,
			})
	}

	return result
}

func (self *DbManager) getColumnInfo(tableName string) []common.ColumnInfo {
	query := self.operator.GetColumnInfo(tableName)
	rows, err := self.db.Query(query)
	if err != nil {
		panic("Failed to query columns info")
	}

	defer rows.Close()

	var result []common.ColumnInfo
	for rows.Next() {
		column := common.ColumnInfo{}

		rows.Scan(&column.ColumnName, &column.ColumnType, &column.Description)

		result = append(result, column)
	}
	return result
}
