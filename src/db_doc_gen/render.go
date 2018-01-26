package db_doc_gen

import (
	"text/template"
	"os"
	"bufio"
	"fmt"
)

func RenderTemplate(tableinfos []TableInfo, cfg Config) {

	tmpl := template.New("tmpl")

	templatePath := cfg.TemplatePath
	outPath := cfg.OutPath

	var err error
	if IsBlank(templatePath) {
		tmpl, err = tmpl.Parse(defaultTmpl)
	} else {
		tmpl, err = tmpl.ParseFiles(templatePath)
	}

	if err != nil {
		panic("Failed to parse template file")
	}
	outFile, _ := os.Create(outPath)
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	tmpl.Execute(writer,
		map[string]interface{}{
			"tables": tableinfos,
			"schema": cfg.Dbinfo.Schema})

	writer.Flush()
	fmt.Printf("save result to %s \n", outPath)
}

const defaultTmpl = `
{{- .schema}} Document
{{range .tables -}}
# {{.Name}}
|column|type|description|
| ------| ------ | ------ |
{{- range .Columns}}
|{{.Name}}|{{.ColumnType}}|{{.Description -}}|
{{- end}}

{{end}}`
