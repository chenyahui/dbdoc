package render

import "github.com/chenyahui/dbdoc/dbdoc/common"

type RenderFunc = func(tableinfos []common.TableInfo, cfg common.Config)

var registry = map[string]RenderFunc{}

func init() {
	registry["word"] = renderWord
	registry["plain"] = renderPlain
}

func GetRenderFunc(name string) RenderFunc {
	return registry[name]
}
