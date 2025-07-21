package search

import (
	"strings"
	"rules-explorer/internal/core/types"
)

type Filter struct {
	query string
}

func NewFilter() *Filter {
	return &Filter{}
}

func (f *Filter) SetQuery(query string) {
	f.query = strings.ToLower(query)
}

func (f *Filter) Match(file types.FileItem) bool {
	if f.query == "" {
		return true
	}
	
	return strings.Contains(strings.ToLower(file.Path), f.query) ||
		strings.Contains(strings.ToLower(file.Content), f.query)
}

func (f *Filter) FilterFiles(files []types.FileItem) []types.FileItem {
	if f.query == "" {
		return files
	}
	
	filtered := make([]types.FileItem, 0)
	for _, file := range files {
		if f.Match(file) {
			filtered = append(filtered, file)
		}
	}
	
	return filtered
}