package ui

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"rules-explorer/internal/file"
)

type App struct {
	app           *tview.Application
	layout        *tview.Flex
	mainLayout    *tview.Flex
	statusBar     *tview.TextView
	
	// Top row panes
	searchInput   *tview.InputField
	fileList      *tview.List
	fileDetails   *tview.TextView
	
	// Bottom row panes
	fileStats     *tview.TextView
	helpPanel     *tview.TextView
	
	// Content preview (spans bottom)
	preview       *tview.TextView
	
	explorer      *file.Explorer
	filteredFiles []file.Item
	currentFocus  int // 0: search, 1: fileList, 2: details, 3: stats, 4: help, 5: preview
}

func NewApp() *App {
	return &App{
		app:          tview.NewApplication(),
		explorer:     file.NewExplorer(),
		currentFocus: 0,
	}
}

func (a *App) Initialize() error {
	a.setupSearchInput()
	a.setupFileList()
	a.setupFileDetails()
	a.setupFileStats()
	a.setupHelpPanel()
	a.setupPreview()
	a.setupStatusBar()
	a.setupLayout()
	a.setupKeybindings()
	
	if err := a.explorer.LoadFiles(); err != nil {
		return err
	}
	
	a.updateFileList("")
	a.updateStatusBar()
	return nil
}

func (a *App) setupSearchInput() {
	a.searchInput = tview.NewInputField().
		SetFieldWidth(0).
		SetPlaceholder("Type to filter files...").
		SetChangedFunc(a.onSearchChanged).
		SetFieldBackgroundColor(tcell.ColorNone).
    SetFieldTextColor(tcell.ColorWhite).
		SetPlaceholderStyle(tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorGray))
	
	a.searchInput.SetBorder(true).
		SetTitle("[yellow]🔍 Search Rules & Config Files[-]").
		SetTitleAlign(tview.AlignLeft).
		SetBackgroundColor(tcell.ColorDefault)
}

func (a *App) setupFileList() {
	a.fileList = tview.NewList().
		ShowSecondaryText(true).
		SetSelectedFunc(a.onFileSelected).
		SetChangedFunc(a.onFileListChanged)
	
	a.fileList.SetBorder(true).
		SetTitle("[aqua]📁 Files[-]").
		SetTitleAlign(tview.AlignLeft)
}

func (a *App) setupFileDetails() {
	a.fileDetails = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetScrollable(true)
	
	a.fileDetails.SetBorder(true).
		SetTitle("[green]📄 File Details[-]").
		SetTitleAlign(tview.AlignLeft)
}

func (a *App) setupFileStats() {
	a.fileStats = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)
	
	a.fileStats.SetBorder(true).
		SetTitle("[magenta]📊 Statistics[-]").
		SetTitleAlign(tview.AlignLeft)
}

func (a *App) setupHelpPanel() {
	a.helpPanel = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)
	
	a.helpPanel.SetBorder(true).
		SetTitle("[blue]❓ Help & Commands[-]").
		SetTitleAlign(tview.AlignLeft)
	
	helpText := `[yellow]Navigation:[-]
[white]Tab[-]       - Switch between panes
[white]↑/↓[-]       - Navigate files
[white]Ctrl+P/N[-]  - Navigate files
[white]Enter[-]     - Select file
[white]Esc[-]       - Exit
[white]Ctrl+C[-]    - Quit

[yellow]File Types:[-]
[red]📋[-] Cursor Rules (.mdc)
[green]📝[-] Claude Config (CLAUDE.md)
[blue]⚙️[-]  Config (.claude/*)`
	a.helpPanel.SetText(helpText)
}

func (a *App) setupPreview() {
	a.preview = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetScrollable(true)
	
	a.preview.SetBorder(true).
		SetTitle("[white]📖 Content Preview[-]").
		SetTitleAlign(tview.AlignLeft)
}

func (a *App) setupStatusBar() {
	a.statusBar = tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(false).
		SetWordWrap(false)
	
	a.statusBar.SetBackgroundColor(tcell.ColorDarkSlateGray)
}

func (a *App) setupLayout() {
	// Left side: Search + File List
	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(a.searchInput, 3, 0, false).
		AddItem(a.fileList, 0, 1, false)
	
	// Main content area: Left (Files) | Right (Preview)
	mainContent := tview.NewFlex().
		AddItem(leftPanel, 0, 1, false).
		AddItem(a.preview, 0, 2, false)
	
	// Bottom info panel: File Details | Stats | Help
	bottomInfo := tview.NewFlex().
		AddItem(a.fileDetails, 0, 1, false).
		AddItem(a.fileStats, 0, 1, false).
		AddItem(a.helpPanel, 0, 1, false)
	
	// Main layout: Top (Files + Preview) | Bottom (Info panels)
	a.mainLayout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainContent, 0, 3, false).
		AddItem(bottomInfo, 0, 1, false)
	
	// Overall layout with status bar
	a.layout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(a.mainLayout, 0, 1, true).
		AddItem(a.statusBar, 1, 0, false)
}

func (a *App) setupKeybindings() {
	a.app.SetInputCapture(a.handleGlobalKeys)
}

func (a *App) onSearchChanged(text string) {
	a.updateFileList(text)
}

func (a *App) updateFileList(filter string) {
	a.fileList.Clear()
	a.filteredFiles = a.explorer.FilterFiles(filter)

	for _, file := range a.filteredFiles {
		icon := a.getFileIcon(file.Path)
		shortPath := a.getShortPath(file.Path)
		fileType := a.getFileType(file.Path)
		a.fileList.AddItem(fmt.Sprintf("%s %s", icon, shortPath), 
			fmt.Sprintf("[gray]%s[-]", fileType), 0, nil)
	}

	if len(a.filteredFiles) > 0 {
		a.fileList.SetCurrentItem(0)
		a.updateFileDetails(a.filteredFiles[0])
		a.updateFileStats()
		a.preview.SetText(a.filteredFiles[0].Content)
	} else {
		allFiles := a.explorer.GetAllFiles()
		a.fileDetails.SetText("[yellow]No files selected[-]")
		a.preview.SetText(fmt.Sprintf("[red]No files found[-]\n\n[white]Total files loaded: %d\nFilter: '%s'[-]", len(allFiles), filter))
	}
	a.updateStatusBar()
}

func (a *App) onFileSelected(index int, mainText string, secondaryText string, shortcut rune) {
	if index < len(a.filteredFiles) {
		a.updateFileDetails(a.filteredFiles[index])
		a.updateFileStats()
		a.preview.Clear()
		a.preview.SetText(a.filteredFiles[index].Content)
		a.updateStatusBar()
	}
}

func (a *App) onFileListChanged(index int, mainText string, secondaryText string, shortcut rune) {
	if index >= 0 && index < len(a.filteredFiles) {
		a.updateFileDetails(a.filteredFiles[index])
		a.updateFileStats()
		a.preview.Clear()
		a.preview.SetText(a.filteredFiles[index].Content)
		a.updateStatusBar()
	}
}

func (a *App) getFileIcon(path string) string {
	if strings.HasSuffix(path, ".mdc") {
		return "[red]📋[-]"
	}
	if filepath.Base(path) == "CLAUDE.md" {
		return "[green]📝[-]"
	}
	if strings.HasPrefix(path, ".claude/") {
		return "[blue]⚙️[-]"
	}
	return "[white]📄[-]"
}

func (a *App) getShortPath(path string) string {
	if len(path) > 40 {
		return "..." + path[len(path)-37:]
	}
	return path
}

func (a *App) getFileType(path string) string {
	if strings.HasSuffix(path, ".mdc") {
		return "Cursor Rule"
	}
	if filepath.Base(path) == "CLAUDE.md" {
		return "Claude Config"
	}
	if strings.HasPrefix(path, ".claude/") {
		return "Configuration"
	}
	return "Unknown"
}

func (a *App) updateFileDetails(file file.Item) {
	lines := strings.Split(file.Content, "\n")
	size := len(file.Content)
	lineCount := len(lines)
	
	var sizeStr string
	if size < 1024 {
		sizeStr = fmt.Sprintf("%d B", size)
	} else if size < 1024*1024 {
		sizeStr = fmt.Sprintf("%.1f KB", float64(size)/1024)
	} else {
		sizeStr = fmt.Sprintf("%.1f MB", float64(size)/(1024*1024))
	}
	
	fileType := a.getFileType(file.Path)
	icon := a.getFileIcon(file.Path)
	
	details := fmt.Sprintf(`%s [white]%s[-]

[yellow]Path:[-] %s
[yellow]Type:[-] %s
[yellow]Size:[-] %s
[yellow]Lines:[-] %d

[yellow]Content Preview:[-]
[gray]%s[-]`, 
		icon, filepath.Base(file.Path),
		file.Path,
		fileType,
		sizeStr,
		lineCount,
		a.getContentPreview(file.Content))
	
	a.fileDetails.SetText(details)
}

func (a *App) getContentPreview(content string) string {
	lines := strings.Split(content, "\n")
	previewLines := []string{}
	
	for i, line := range lines {
		if i >= 5 { // Show first 5 lines
			previewLines = append(previewLines, "...")
			break
		}
		if len(line) > 50 {
			line = line[:47] + "..."
		}
		previewLines = append(previewLines, line)
	}
	
	return strings.Join(previewLines, "\n")
}

func (a *App) updateFileStats() {
	allFiles := a.explorer.GetAllFiles()
	filteredCount := len(a.filteredFiles)
	totalCount := len(allFiles)
	
	// Count by type
	cursorRules := 0
	claudeConfigs := 0
	configFiles := 0
	
	for _, file := range allFiles {
		if strings.HasSuffix(file.Path, ".mdc") {
			cursorRules++
		} else if filepath.Base(file.Path) == "CLAUDE.md" {
			claudeConfigs++
		} else if strings.HasPrefix(file.Path, ".claude/") {
			configFiles++
		}
	}
	
	stats := fmt.Sprintf(`[yellow]Total Files:[-] %d
[yellow]Filtered:[-] %d

[yellow]By Type:[-]
[red]📋[-] Cursor Rules: %d
[green]📝[-] Claude Configs: %d
[blue]⚙️[-] Config Files: %d

[yellow]Timestamp:[-]
%s`,
		totalCount,
		filteredCount,
		cursorRules,
		claudeConfigs,
		configFiles,
		time.Now().Format("15:04:05"))
	
	a.fileStats.SetText(stats)
}

func (a *App) updateStatusBar() {
	currentFile := "None"
	if len(a.filteredFiles) > 0 {
		currentIndex := a.fileList.GetCurrentItem()
		if currentIndex >= 0 && currentIndex < len(a.filteredFiles) {
			currentFile = filepath.Base(a.filteredFiles[currentIndex].Path)
		}
	}
	
	statusText := fmt.Sprintf(" [yellow]Rules Explorer[-] | Files: %d/%d | Current: [aqua]%s[-] | [white]Tab[-]: Switch Panes | [white]Esc[-]: Exit",
		len(a.filteredFiles), len(a.explorer.GetAllFiles()), currentFile)
	
	a.statusBar.SetText(statusText)
}

func (a *App) handleGlobalKeys(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyTab:
		a.switchFocus()
		return nil
	case tcell.KeyCtrlC, tcell.KeyEscape:
		a.app.Stop()
		return nil
	case tcell.KeyEnter:
		if a.app.GetFocus() == a.fileList {
			currentIndex := a.fileList.GetCurrentItem()
			a.onFileSelected(currentIndex, "", "", 0)
		}
		return nil
	}
	
	// Handle navigation in file list
	if a.app.GetFocus() == a.fileList {
		switch event.Key() {
		case tcell.KeyDown, tcell.KeyCtrlN:
			current := a.fileList.GetCurrentItem()
			if current < a.fileList.GetItemCount()-1 {
				a.fileList.SetCurrentItem(current + 1)
			}
			return nil
		case tcell.KeyUp, tcell.KeyCtrlP:
			current := a.fileList.GetCurrentItem()
			if current > 0 {
				a.fileList.SetCurrentItem(current - 1)
			}
			return nil
		}
	}
	
	return event
}

func (a *App) switchFocus() {
	switch a.currentFocus {
	case 0: // search -> fileList
		a.app.SetFocus(a.fileList)
		a.currentFocus = 1
		a.updateBorderColors()
	case 1: // fileList -> search
		a.app.SetFocus(a.searchInput)
		a.currentFocus = 0
		a.updateBorderColors()
	}
}

func (a *App) updateBorderColors() {
	// Set all borders to the unfocused color #C8D3F5
	unfocusedColor := tcell.NewHexColor(0xC8D3F5)
	
	a.searchInput.SetBorderColor(unfocusedColor)
	a.fileList.SetBorderColor(unfocusedColor)
	a.fileDetails.SetBorderColor(unfocusedColor)
	a.fileStats.SetBorderColor(unfocusedColor)
	a.helpPanel.SetBorderColor(unfocusedColor)
	a.preview.SetBorderColor(unfocusedColor)
	
	// Highlight current focus with the original colors
	switch a.currentFocus {
	case 0:
		a.searchInput.SetBorderColor(tcell.ColorYellow)
	case 1:
		a.fileList.SetBorderColor(tcell.ColorTeal)
	}
}

func (a *App) Run() error {
	a.updateBorderColors()
	return a.app.SetRoot(a.layout, true).SetFocus(a.searchInput).Run()
}
