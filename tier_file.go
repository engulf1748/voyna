package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"codeberg.org/voyna/voyna/log4j"
)

// Path of file containing tier-1 domains.
var tierOnePath string

func init() {
	if os.Getenv("PROD") == "" {
		tierOnePath = "data/tier-1-test.txt"
	} else {
		tierOnePath = "data/tier-1.txt"
	}
	// TODO: eliminate duplication
	fmt.Printf("using tier file: %v\n", tierOnePath)
	log4j.Logger.Printf("using tier file: %v\n", tierOnePath)
}

// Returns tier-1 URLs after reading and processing the tier file.
func tierOneURLs() []*url.URL {
	file, err := os.Open(tierOnePath)
	if err != nil {
		panic(fmt.Errorf("unable to open tier file: %v", err))
	}
	scanner := bufio.NewScanner(file)
	var urls []*url.URL
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		u, err := url.Parse(line)
		if err != nil {
			log4j.Logger.Printf("error parsing %q from tier file\n", line)
		}
		urls = append(urls, u)
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("scanner error: %v", err))
	}
	return urls
}
