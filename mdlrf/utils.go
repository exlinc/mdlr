package mdlrf

import (
	"os"
	"path/filepath"
)

func getMdlrFilePathForCmd() (string, error) {
	path := os.Getenv("MDLR_FILE")
	if path != "" {
		return filepath.Abs(path)
	}
	return filepath.Abs("mdlr.yml")
}
