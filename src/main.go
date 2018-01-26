package main

import (
	"db_doc_gen"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %s \n", err)
		}
	}()

	cfg := db_doc_gen.ParseCmd()

	status, msg := db_doc_gen.CheckConfig(&cfg)
	if !status {
		panic(msg)
	}

	var db = db_doc_gen.DbManager{}
	defer db.Close()

	db.Connect(cfg)
	db_doc_gen.RenderTemplate(db.GetTablesInfo(), cfg)
}
