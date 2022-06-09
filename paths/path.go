// Contains default paths you might want to use. This package does not create
// the directories or files it names, however, so please create them yourself.
package paths

import (
	"os"
	"path/filepath"
)

var (
	homeDir       string
	relStorageDir = filepath.Join(".local", "share", "voyna")
	StorageDir    string
	CorpusDir     string
	LogsDir       string
	EasyLogPath   string
)

const (
	corpusDir   = "corpus"
	logsDir     = "logs"
	easyLogFile = "log"
)

func init() {
	var err error
	homeDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	StorageDir = join(homeDir, relStorageDir)
	CorpusDir = join(StorageDir, corpusDir)
	LogsDir = join(StorageDir, logsDir)
	EasyLogPath = join(LogsDir, easyLogFile)
}

func join(dirs ...string) string {
	return filepath.Join(dirs...)
}
