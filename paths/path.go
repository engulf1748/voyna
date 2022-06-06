package paths

import (
	"os"
	"path/filepath"
)

var HomeDir string
var RelDataPath = filepath.Join(".local", "share", "voyna")

const (
	CrawlDirName = "crawler"
)

func init() {
	var err error
	HomeDir, err = os.UserHomeDir()
	if err != nil {
		// TODO
	}
}

func DataDir(subdir ...string) string {
	return filepath.Join(HomeDir, RelDataPath, filepath.Join(subdir...))
}

func CrawlDir() string {
	return filepath.Join(DataDir(CrawlDirName))
}
