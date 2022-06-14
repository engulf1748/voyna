package main

import (
	"sync"

	"codeberg.org/voyna/voyna/processor"
)

func main() {
	urls := tierOneURLs()
	go processor.Process(urls)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait() // waits indefinitely
}
