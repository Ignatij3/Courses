package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "INFO: ", log.Lshortfile)

		infof = func(calldepth int, info string) {
			logger.Output(calldepth, info)
		}
	)

	infof(1, "Hello world")		//  INFO: sample_log.go:15: Hello world
	infof(2, "Hello world")		//  INFO: sample_log.go:20: Hello world
	infof(3, "Hello world")		//  INFO: proc.go:203: Hello world
	infof(4, "Hello world")		//  INFO: asm_amd64.s:1357: Hello world
	infof(5, "Hello world")		//  INFO: ???:0: Hello world

	fmt.Print(&buf)
}




