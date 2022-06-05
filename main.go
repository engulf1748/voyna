package main

import (
	"codeberg.org/voyna/voyna/processor"
	"codeberg.org/voyna/voyna/server"
)

func main() {
	domains := readDomains()
	go processor.Process(domains)
	server.Start()
}
