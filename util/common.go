package util

import (
	"os"
)

func ReadFiles(pathFile string) ([]byte, error) {
	return os.ReadFile(pathFile)
}
