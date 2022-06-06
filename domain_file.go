package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"codeberg.org/voyna/voyna/log4j"
)

// Path of domain-list file relative to this file.
var domainFilePath string

func init() {
	if os.Getenv("DEV") == "" {
		domainFilePath = "data/tier-1-test.txt"
	} else {
		domainFilePath = "data/tier-1.txt"
	}
	// TODO: eliminate duplication
	fmt.Printf("using domains from: %v\n", domainFilePath)
	log4j.Logger.Printf("using domains from: %v\n", domainFilePath)
}

// Scans the domains file line-by-line and returns
// a slice of all non-empty and all non-commented lines.
func readDomains() []string {
	file, err := os.Open(domainFilePath)
	if err != nil {
		panic(fmt.Errorf("unable to open domains file: %v", err))
	}
	scanner := bufio.NewScanner(file)
	var domains []string
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain == "" || strings.HasPrefix(domain, "#") {
			continue
		}
		domains = append(domains, domain)
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("scanner error: %v", err))
	}
	return domains
}
