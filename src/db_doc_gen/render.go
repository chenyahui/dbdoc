package db_doc_gen

import (
	"text/template"
	"os"
	"bufio"
)

func RenderTemplate(tableinfos []TableInfo, template_path string, out_path string) {

	tmpl := template.Must(template.ParseFiles(template_path))

	outFile, _ := os.Create(out_path)
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	tmpl.Execute(writer, tableinfos)

	writer.Flush()
}