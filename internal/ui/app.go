package ui

import (
	"fmt"

	"github.com/rivo/tview"

	"rules-explorer/internal/file"
)

type App struct {
	app           *tview.Application
	layout        *tview.Flex
	searchInput   *tview.InputField
	fileList      *tview.List
	preview       *tview.TextView
	explorer      *file.Explorer
	filteredFiles []file.Item
}

func NewApp() *App {
	return &App{
		app:      tview.NewApplication(),
		explorer: file.NewExplorer(),
	}
}

func (a *App) Initialize() error {
	a.setupSearchInput()
	a.setupFileList()
	a.setupPreview()
	a.setupLayout()
	a.setupKeybindings()
	
	if err := a.explorer.LoadFiles(); err != nil {
		return err
	}
	
	a.updateFileList("")
	return nil
}

func (a *App) setupSearchInput() {
	a.searchInput = tview.NewInputField().
		SetFieldWidth(0).
		SetDoneFunc(HandleSearchEscape(a.app)).
		SetChangedFunc(a.onSearchChanged)
	
	a.searchInput.SetBorder(true).SetTitle("Search")
}

func (a *App) setupFileList() {
	a.fileList = tview.NewList().
		ShowSecondaryText(false).
		SetSelectedFunc(a.onFileSelected).
		SetChangedFunc(a.onFileListChanged)
	
	a.fileList.SetBorder(true).SetTitle("Files")
	a.fileList.SetInputCapture(HandleFileListKeys(a.fileList))
}

func (a *App) setupPreview() {
	a.preview = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)
	
	a.preview.SetBorder(true).SetTitle("Preview")
}

func (a *App) setupLayout() {
	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(a.searchInput, 3, 0, true).
		AddItem(a.fileList, 0, 1, false)

	a.layout = tview.NewFlex().
		AddItem(leftPanel, 0, 1, true).
		AddItem(a.preview, 0, 1, false)
}

func (a *App) setupKeybindings() {
	a.app.SetInputCapture(HandleGlobalKeys(a.app, a.searchInput, a.fileList))
}

func (a *App) onSearchChanged(text string) {
	a.updateFileList(text)
}

func (a *App) updateFileList(filter string) {
	a.fileList.Clear()
	a.filteredFiles = a.explorer.FilterFiles(filter)

	for _, file := range a.filteredFiles {
		a.fileList.AddItem(file.Path, "", 0, nil)
	}

	if len(a.filteredFiles) > 0 {
		a.fileList.SetCurrentItem(0)
		a.preview.SetText(a.filteredFiles[0].Content)
	} else {
		allFiles := a.explorer.GetAllFiles()
		a.preview.SetText(fmt.Sprintf("No files found.\nTotal files loaded: %d\nFilter: '%s'", len(allFiles), filter))
	}
}

func (a *App) onFileSelected(index int, mainText string, secondaryText string, shortcut rune) {
	if index < len(a.filteredFiles) {
		a.preview.Clear()
		a.preview.SetText(a.filteredFiles[index].Content)
	}
}

func (a *App) onFileListChanged(index int, mainText string, secondaryText string, shortcut rune) {
	if index >= 0 && index < len(a.filteredFiles) {
		a.preview.Clear()
		a.preview.SetText(a.filteredFiles[index].Content)
	}
}

func (a *App) Run() error {
	return a.app.SetRoot(a.layout, true).SetFocus(a.searchInput).Run()
}