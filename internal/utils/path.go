package utils

import (
	"path/filepath"
)

func GetBaseName(path string) string {
	return filepath.Base(path)
}

func GetShortPath(path string, maxLength int) string {
	if len(path) > maxLength {
		return "..." + path[len(path)-(maxLength-3):]
	}
	return path
}