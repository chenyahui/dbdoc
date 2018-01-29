package main

import (
	"dbdoc"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %s \n", err)
		}
	}()

	cfg := dbdoc.ParseCmd()

	status, msg := dbdoc.CheckConfig(&cfg)
	if !status {
		panic(msg)
	}

	var db = dbdoc.DbManager{}
	defer db.Close()

	db.Connect(cfg)
	dbdoc.RenderTemplate(db.GetTablesInfo(), cfg)
}
