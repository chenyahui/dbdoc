package dbdoc

import (
	"baliance.com/gooxml/document"
)

func RenderToWord(tableinfos []TableInfo, cfg Config) {
	doc := document.New()

	for _, tableInfo := range tableinfos {
		table := doc.AddTable()
		addTableName(&table, tableInfo.TableName)
		addRow(&table, ColumnInfo{"column", "type", "description"})
		for _, columnInfo := range tableInfo.Columns {
			addRow(&table, columnInfo)
		}
		doc.AddParagraph()
	}
	doc.SaveToFile(cfg.OutPath)
}
func addTableName(table *document.Table, tableName string) {
	row := table.AddRow()
	cell := row.AddCell()
	cell.Properties().SetColumnSpan(3)
	cell.AddParagraph().AddRun().AddText(tableName)
}
func addRow(table *document.Table, columnInfo ColumnInfo) {
	row := table.AddRow()
	row.AddCell().AddParagraph().AddRun().AddText(columnInfo.ColumnName)
	row.AddCell().AddParagraph().AddRun().AddText(columnInfo.ColumnType)
	row.AddCell().AddParagraph().AddRun().AddText(columnInfo.Description)
}
