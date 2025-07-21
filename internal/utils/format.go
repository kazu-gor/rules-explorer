package utils

import (
	"fmt"
	"strings"
)

func FormatFileSize(size int) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.1f KB", float64(size)/1024)
	} else {
		return fmt.Sprintf("%.1f MB", float64(size)/(1024*1024))
	}
}

func GetContentPreview(content string, maxLines int, maxLineLength int) string {
	lines := strings.Split(content, "\n")
	previewLines := []string{}
	
	for i, line := range lines {
		if i >= maxLines {
			previewLines = append(previewLines, "...")
			break
		}
		if len(line) > maxLineLength {
			line = line[:maxLineLength-3] + "..."
		}
		previewLines = append(previewLines, line)
	}
	
	return strings.Join(previewLines, "\n")
}

func CountLines(content string) int {
	return len(strings.Split(content, "\n"))
}