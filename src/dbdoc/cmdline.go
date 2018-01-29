package dbdoc

import "flag"

func ParseCmd() Config {
	configPath := flag.String("c", "", "config file")
	flag.Parse()

	if IsBlank(*configPath) {
		panic("no config file")
	}

	return ParseConfigFile(*configPath)
}
