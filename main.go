package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type FileItem struct {
	Path    string
	Content string
}

var (
	app         *tview.Application
	layout      *tview.Flex
	searchInput *tview.InputField
	fileList    *tview.List
	preview     *tview.TextView
	allFiles    []FileItem
	filteredFiles []FileItem
)

func main() {
	app = tview.NewApplication()

	searchInput = tview.NewInputField().
		SetFieldWidth(0).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				app.Stop()
			}
		}).
		SetChangedFunc(onSearchChanged)

	fileList = tview.NewList().
		ShowSecondaryText(false).
		SetSelectedFunc(onFileSelected).
		SetChangedFunc(onFileListChanged)
	fileList.SetBorder(true).SetTitle("Files")
	
	// Add Ctrl+N/Ctrl+P navigation to fileList
	fileList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlN:
			// Move down
			current := fileList.GetCurrentItem()
			if current < fileList.GetItemCount()-1 {
				fileList.SetCurrentItem(current + 1)
			}
			return nil
		case tcell.KeyCtrlP:
			// Move up
			current := fileList.GetCurrentItem()
			if current > 0 {
				fileList.SetCurrentItem(current - 1)
			}
			return nil
		}
		return event
	})

	preview = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)
	preview.SetBorder(true).SetTitle("Preview")

	searchInput.SetBorder(true).SetTitle("Search")

	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(searchInput, 3, 0, true).
		AddItem(fileList, 0, 1, false)

	layout = tview.NewFlex().
		AddItem(leftPanel, 0, 1, true).
		AddItem(preview, 0, 1, false)

	loadFiles()
	updateFileList("")

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			if app.GetFocus() == searchInput {
				app.SetFocus(fileList)
			} else {
				app.SetFocus(searchInput)
			}
			return nil
		case tcell.KeyCtrlC:
			app.Stop()
			return nil
		}
		
		// When fileList has focus, ensure arrow keys work
		if app.GetFocus() == fileList {
			switch event.Key() {
			case tcell.KeyDown:
				current := fileList.GetCurrentItem()
				if current < fileList.GetItemCount()-1 {
					fileList.SetCurrentItem(current + 1)
				}
				return nil
			case tcell.KeyUp:
				current := fileList.GetCurrentItem()
				if current > 0 {
					fileList.SetCurrentItem(current - 1)
				}
				return nil
			}
		}
		
		return event
	})

	if err := app.SetRoot(layout, true).SetFocus(searchInput).Run(); err != nil {
		panic(err)
	}
}

func loadFiles() {
	allFiles = []FileItem{}
	cwd, err := os.Getwd()
	if err != nil {
		return
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

			allFiles = append(allFiles, FileItem{
				Path:    relPath,
				Content: string(content),
			})
		}

		return nil
	})

	if err != nil {
		return
	}
}

func onSearchChanged(text string) {
	updateFileList(text)
}

func updateFileList(filter string) {
	fileList.Clear()
	filteredFiles = []FileItem{}

	for _, file := range allFiles {
		if filter == "" || strings.Contains(strings.ToLower(file.Path), strings.ToLower(filter)) ||
			strings.Contains(strings.ToLower(file.Content), strings.ToLower(filter)) {
			filteredFiles = append(filteredFiles, file)
			fileList.AddItem(file.Path, "", 0, nil)
		}
	}

	if len(filteredFiles) > 0 {
		fileList.SetCurrentItem(0)
		preview.SetText(filteredFiles[0].Content)
	} else {
		preview.SetText(fmt.Sprintf("No files found.\nTotal files loaded: %d\nFilter: '%s'", len(allFiles), filter))
	}
}

func onFileSelected(index int, mainText string, secondaryText string, shortcut rune) {
	if index < len(filteredFiles) {
		preview.Clear()
		preview.SetText(filteredFiles[index].Content)
	}
}

func onFileListChanged(index int, mainText string, secondaryText string, shortcut rune) {
	if index >= 0 && index < len(filteredFiles) {
		preview.Clear()
		preview.SetText(filteredFiles[index].Content)
	}
}
