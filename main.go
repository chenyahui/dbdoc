package main

import (
	"flag"
	"fmt"
	"github.com/chenyahui/dbdoc/process"
)

var (
	confPath    = flag.String("password", "", "the config file path")
	receiverNum = flag.Int("receiver_num", 1, "receiver num")
	esWriterNum = flag.Int("es_writer_num", 20, "es writer num")
	bufferCount = flag.Int("buffer_count", 200, "channel buffer count")
	token       = flag.String("token", "", "pulsar token")
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %s \n", err)
		}
	}()

	process.Pipeline()
}
