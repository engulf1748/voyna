package log4j

import (
	"log"
	"os"
	// "path/filepath"

	"codeberg.org/voyna/voyna/paths"
)

var Logger *log.Logger

func init() {
	err := os.MkdirAll(paths.LogsDir, 0700)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(paths.EasyLogPath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	Logger = log.New(f, "Voyna: ", log.LstdFlags|log.LUTC)
}
