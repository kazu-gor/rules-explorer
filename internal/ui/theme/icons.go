package theme

import (
	"path/filepath"
	"strings"
	"rules-explorer/internal/core/types"
)

func DetermineFileType(path string) types.FileType {
	if strings.HasSuffix(path, ".mdc") {
		return types.CursorRule
	}
	if filepath.Base(path) == "CLAUDE.md" {
		return types.ClaudeConfig
	}
	if strings.HasPrefix(path, ".claude/") {
		return types.ConfigFile
	}
	return types.Unknown
}

func GetFileIcon(path string, icons types.IconSet) string {
	fileType := DetermineFileType(path)
	return GetFileTypeIcon(fileType, icons)
}

func FormatPath(path string) string {
	if len(path) > 80 {
		return "..." + path[len(path)-77:]
	}
	return path
}