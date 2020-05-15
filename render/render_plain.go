package render

import (
	"bufio"
	"fmt"
	"github.com/chenyahui/dbdoc/common"
	"os"
	"text/template"
)

func renderPlain(tableinfos []common.TableInfo, cfg common.Config) {
	tmpl := template.New("tmpl")

	templatePath := cfg.TemplatePath
	outPath := cfg.OutPath

	var err error
	if common.IsBlank(templatePath) {
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
			"schema": cfg.DbInfo.Schema})

	writer.Flush()
	fmt.Printf("save result to %s \n", outPath)
}

const defaultTmpl = `
{{- .schema}} Document
{{range .tables -}}
# {{.TableName}}
|column|type|description|
| ------| ------ | ------ |
{{- range .Columns}}
|{{.ColumnName}}|{{.ColumnType}}|{{.Description -}}|
{{- end}}

{{end}}`
