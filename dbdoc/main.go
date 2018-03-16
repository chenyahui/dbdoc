package dbdoc

import (
	"github.com/chenyahui/dbdoc/dbdoc/process"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %s \n", err)
		}
	}()

	process.Pipeline()
}
