# Database to Markdown

## How to use

```
db_doc_generator -c config.json
```

config file
```
{
  "db_info": {
    "db_type": "mysql",  // required
    "ip_port": "127.0.0.1:3306", // required
    "username": "root", // required
    "password": "", // required
    "schema": "cms" // required
  },
  "includes": [  // optional
  ],
  "excludes": [  // optional
  ],
  "template_path": "", // optional
  "out_path": ""  // optional
}
```

|Property|Description|Required|
| ------| ------ |------ |
|db_type|database driver name|Required|
|ip_port|database's host and port|Required|
|username|username|Required|
|password|password|Required|
|schema|the schema you want to read|Required|
|includes|the tables you want to generate. Default: all tables|Optional|
|excludes|the tables you dont't want to generate.|Optional|
|template_path|markdown's layout|Optional|
|out_path|the path you want to save. Default: `schema`_doc.md|Optional|

# template
You can use your own markdown template by set the template_path.
The tool also offer a default one for you.

```
{{- .schema}} Document
{{range .tables -}}
# {{.TableName}}
|column|type|description|
| ------| ------ | ------ |
{{- range .Columns}}
|{{.ColumnName}}|{{.ColumnType}}|{{.Description -}}|
{{- end}}

{{end}}
```

The template use the golang `text/template`, you can read the grammar at [here](https://golang.org/pkg/text/template/)


