package main

import (
	"sync"

	"codeberg.org/voyna/voyna/processor"
	"codeberg.org/voyna/voyna/server"
)

func main() {
	urls := tierOneURLs()
	go processor.Process(urls)

	server.Start()

	// TODO: remove the following
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait() // waits indefinitely
}
