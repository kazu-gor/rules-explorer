package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Item struct {
	Path    string
	Content string
}

type Explorer struct {
	allFiles []Item
}

func NewExplorer() *Explorer {
	return &Explorer{
		allFiles: []Item{},
	}
}

func (e *Explorer) LoadFiles() error {
	e.allFiles = []Item{}
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	err = filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(cwd, path)
		if err != nil {
			return nil
		}

		// Check if file matches our patterns
		matched := false
		
		// Match .cursor/rules/*.mdc
		if strings.HasPrefix(relPath, ".cursor/rules/") && strings.HasSuffix(relPath, ".mdc") {
			matched = true
		}
		
		// Match CLAUDE.md anywhere
		if filepath.Base(relPath) == "CLAUDE.md" {
			matched = true
		}
		
		// Match .claude/* (direct children only)
		if strings.HasPrefix(relPath, ".claude/") && !strings.Contains(relPath[8:], "/") {
			matched = true
		}

		if matched {
			content, err := os.ReadFile(path)
			if err != nil {
				content = []byte(fmt.Sprintf("Error reading file: %v", err))
			}

			e.allFiles = append(e.allFiles, Item{
				Path:    relPath,
				Content: string(content),
			})
		}

		return nil
	})

	return err
}

func (e *Explorer) FilterFiles(filter string) []Item {
	if filter == "" {
		return e.allFiles
	}

	filtered := []Item{}
	lowerFilter := strings.ToLower(filter)
	
	for _, file := range e.allFiles {
		if strings.Contains(strings.ToLower(file.Path), lowerFilter) ||
			strings.Contains(strings.ToLower(file.Content), lowerFilter) {
			filtered = append(filtered, file)
		}
	}
	
	return filtered
}

func (e *Explorer) GetAllFiles() []Item {
	return e.allFiles
}
