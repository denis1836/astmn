package seeker

import (
	"fmt"
	"os"
	"path/filepath"

	"astmn/internal/log"
)

func SeekAstmnDir(startDir string) (string, error) {
	if startDir == "" {
		return "", fmt.Errorf("start dir field is empty")
	}

	absStartDir, err := filepath.Abs(startDir)
	if err != nil {
		return "", err
	}

	currDir := absStartDir
	for {
		astmnPath := filepath.Join(currDir, ".astmn")

		info, err := os.Stat(astmnPath)
		if err == nil && info.IsDir() {
			log.Infof(".astmn project dir located at %s", astmnPath)
			return astmnPath, nil
		}

		parentDir := filepath.Dir(currDir)
		if parentDir == currDir {
			return "", fmt.Errorf("unable to find .astmn project dir")
		}

		currDir = parentDir
	}
}
