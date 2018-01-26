# Database to Markdown

## How to use

```
db_doc_generator -c config.json
```

config file
```json
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