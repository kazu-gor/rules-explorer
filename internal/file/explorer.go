package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"rules-explorer/internal/core/types"
	"rules-explorer/internal/core/search"
)

type Explorer struct {
	allFiles []types.FileItem
	filter   *search.Filter
}

func NewExplorer() *Explorer {
	return &Explorer{
		allFiles: make([]types.FileItem, 0),
		filter:   search.NewFilter(),
	}
}

func (e *Explorer) LoadFiles() error {
	e.allFiles = make([]types.FileItem, 0)
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

		if e.matchesPattern(relPath) {
			content, err := os.ReadFile(path)
			if err != nil {
				content = []byte(fmt.Sprintf("Error reading file: %v", err))
			}

			e.allFiles = append(e.allFiles, types.FileItem{
				Path:    relPath,
				Content: string(content),
			})
		}

		return nil
	})

	return err
}

func (e *Explorer) matchesPattern(relPath string) bool {
	// Match .cursor/rules/*.mdc anywhere in the tree
	if strings.Contains(relPath, ".cursor/rules/") && strings.HasSuffix(relPath, ".mdc") {
		// Ensure the file is directly under .cursor/rules/, not in a subdirectory
		parts := strings.Split(relPath, ".cursor/rules/")
		if len(parts) == 2 && !strings.Contains(parts[1], "/") {
			return true
		}
	}
	
	// Match CLAUDE.md anywhere
	if filepath.Base(relPath) == "CLAUDE.md" {
		return true
	}
	
	// Match .claude/* (direct children only)
	if strings.HasPrefix(relPath, ".claude/") && !strings.Contains(relPath[8:], "/") {
		return true
	}

	return false
}

func (e *Explorer) FilterFiles(filter string) []types.FileItem {
	e.filter.SetQuery(filter)
	return e.filter.FilterFiles(e.allFiles)
}

func (e *Explorer) GetAllFiles() []types.FileItem {
	return e.allFiles
}
