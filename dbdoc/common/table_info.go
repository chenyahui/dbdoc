package common

type ColumnInfo struct {
	ColumnName  string
	ColumnType  string
	Description string
}

type TableInfo struct {
	Columns   []ColumnInfo
	TableName string
}

