package process

import (
	"github.com/chenyahui/dbdoc/dbdoc/cmdline"
	"github.com/chenyahui/dbdoc/dbdoc/common"
	"github.com/chenyahui/dbdoc/dbdoc/render"
)

func Pipeline() {
	cfg := cmdline.ParseCmd()

	if status, msg := common.CheckConfig(&cfg); !status {
		panic(msg)
	}

	var db = DbManager{}
	defer db.Close()

	db.Connect(cfg)
	render.GetRenderFunc(cfg.DocType)(db.GetTablesInfo(), cfg)
}
