package main

import (
	"db_doc_gen"
)

func main() {
	cfg := db_doc_gen.ParseConfigFile("F:\\iopensource\\db_doc_generator\\etc\\default_config.json")
	var db = db_doc_gen.DbManager{}

	err := db.Connect(cfg)
	if (err != nil) {
		panic(err)
	}
	db_doc_gen.RenderTemplate(db.GetTablesInfo(),
		"F:\\iopensource\\db_doc_generator\\etc\\default_template.md",
		"out.md",
	)
}
