package main

import (
	"fmt"
	"runtime"
	//"sync"
	"time"

	"codeberg.org/voyna/voyna/log4j"
	"codeberg.org/voyna/voyna/processor"
	"codeberg.org/voyna/voyna/server"
)

func main() {
	urls := tierOneURLs()
	go processor.Process(urls)

	go server.Start()

	for {
		s := fmt.Sprintf("NumGoroutine: %d\n", runtime.NumGoroutine())
		log4j.Logger.Print(s)
		fmt.Print(s)
		time.Sleep(time.Second)
	}
	// TODO: remove the following
	// 	var wg sync.WaitGroup
	// 	wg.Add(1)
	// 	wg.Wait() // waits indefinitely
}
