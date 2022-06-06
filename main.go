package main

import (
	"sync"

	"codeberg.org/voyna/voyna/processor"
)

func main() {
	domains := readDomains()
	go processor.Process(domains)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait() // waits indefinitely
}
