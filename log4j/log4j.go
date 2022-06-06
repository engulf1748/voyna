package log4j

import (
	"log"
	"os"

	"codeberg.org/voyna/voyna/paths"
)

var Logger *log.Logger

func init() {
	f, err := os.OpenFile(paths.DataDir("log"), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	Logger = log.New(f, "Voyna: ", log.LstdFlags|log.LUTC)
}
