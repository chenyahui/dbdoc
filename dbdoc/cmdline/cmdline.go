package cmdline

import "flag"
import (
	"github.com/chenyahui/dbdoc/dbdoc/common"
)

func ParseCmd() common.Config {
	configPath := flag.String("c", "", "config file")
	flag.Parse()

	if common.IsBlank(*configPath) {
		panic("no config file")
	}

	return common.ParseConfigFile(*configPath)
}
