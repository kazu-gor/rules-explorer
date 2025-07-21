package app

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"rules-explorer/internal/core/types"
	"rules-explorer/internal/file"
	"rules-explorer/internal/ui/input"
	"rules-explorer/internal/ui/layout"
	"rules-explorer/internal/ui/theme"
)

type App struct {
	tvApp         *tview.Application
	config        *Config
	theme         types.Theme
	layoutManager *layout.Manager
	keyHandler    *input.KeyboardHandler
	explorer      types.FileExplorer
	
	// State
	allFiles      []types.FileItem
	filteredFiles []types.FileItem
	currentFile   *types.FileItem
}

func New() *App {
	config := NewConfig()
	appTheme := theme.New()
	config.Theme = appTheme
	
	return &App{
		tvApp:    tview.NewApplication(),
		config:   config,
		theme:    appTheme,
		explorer: file.NewExplorer(),
	}
}

func (a *App) Initialize() error {
	// Load files
	if err := a.explorer.LoadFiles(); err != nil {
		return fmt.Errorf("failed to load files: %w", err)
	}
	
	a.allFiles = a.explorer.GetAllFiles()
	a.filteredFiles = a.allFiles
	
	// Setup UI
	a.layoutManager = layout.NewManager(a.theme)
	a.keyHandler = input.NewKeyboardHandler(a.tvApp)
	
	// Register components with keyboard handler
	components := a.layoutManager.GetComponents()
	for focus, component := range components {
		a.keyHandler.RegisterComponent(focus, component)
	}
	
	// Setup event handlers
	a.setupEventHandlers()
	
	// Setup keyboard handling
	a.tvApp.SetInputCapture(a.keyHandler.HandleGlobalKeys)
	
	// Initial data setup
	a.updateAllComponents()
	
	return nil
}

func (a *App) setupEventHandlers() {
	// Main event handler
	a.keyHandler.SetEventHandler(a.handleEvent)
	
	// Component event handlers
	a.layoutManager.GetSearchComponent().SetEventHandler(a.handleEvent)
	a.layoutManager.GetFileListComponent().SetEventHandler(a.handleEvent)
	a.layoutManager.GetPreviewComponent().SetEventHandler(a.handleEvent)
	a.layoutManager.GetDetailsComponent().SetEventHandler(a.handleEvent)
	a.layoutManager.GetStatsComponent().SetEventHandler(a.handleEvent)
	a.layoutManager.GetHelpComponent().SetEventHandler(a.handleEvent)
	a.layoutManager.GetStatusBarComponent().SetEventHandler(a.handleEvent)
}

func (a *App) handleEvent(event types.Event) {
	switch event.Type {
	case types.EventSearchChanged:
		if searchEvent, ok := event.Data.(types.SearchEvent); ok {
			a.handleSearchChanged(searchEvent.Query)
		}
	case types.EventFileSelected:
		if fileEvent, ok := event.Data.(types.FileEvent); ok {
			a.handleFileSelected(fileEvent.File, fileEvent.Index)
		}
	case types.EventFileChanged:
		if fileEvent, ok := event.Data.(types.FileEvent); ok {
			a.handleFileChanged(fileEvent.File, fileEvent.Index)
		}
	case types.EventFocusChanged:
		if focusEvent, ok := event.Data.(types.FocusEvent); ok {
			a.handleFocusChanged(focusEvent.Focus)
		}
	case types.EventRefresh:
		a.handleRefresh()
	case types.EventQuit:
		a.tvApp.Stop()
	}
}

func (a *App) handleSearchChanged(query string) {
	a.filteredFiles = a.explorer.FilterFiles(query)
	a.layoutManager.GetFileListComponent().Update(a.filteredFiles)
	a.layoutManager.GetStatsComponent().SetFilteredFiles(a.filteredFiles)
	
	// Update preview and details with first file if available
	if len(a.filteredFiles) > 0 {
		a.currentFile = &a.filteredFiles[0]
		a.layoutManager.GetPreviewComponent().Update(*a.currentFile)
		a.layoutManager.GetDetailsComponent().Update(*a.currentFile)
		a.layoutManager.GetStatusBarComponent().Update(*a.currentFile)
	} else {
		a.currentFile = nil
		a.layoutManager.GetPreviewComponent().Update(fmt.Sprintf("[red]No files found[-]\n\n[white]Total files loaded: %d\nFilter: '%s'[-]", len(a.allFiles), query))
		a.layoutManager.GetDetailsComponent().SetNoFileSelected()
		a.layoutManager.GetStatusBarComponent().Update("")
	}
	
	a.layoutManager.GetStatusBarComponent().SetCounts(len(a.filteredFiles), len(a.allFiles))
}

func (a *App) handleFileSelected(file types.FileItem, index int) {
	a.currentFile = &file
	a.layoutManager.GetPreviewComponent().Update(file)
	a.layoutManager.GetDetailsComponent().Update(file)
	a.layoutManager.GetStatusBarComponent().Update(file)
}

func (a *App) handleFileChanged(file types.FileItem, index int) {
	a.currentFile = &file
	a.layoutManager.GetPreviewComponent().Update(file)
	a.layoutManager.GetDetailsComponent().Update(file)
	a.layoutManager.GetStatusBarComponent().Update(file)
}

func (a *App) handleFocusChanged(focus types.Focus) {
	// Focus change is already handled by the keyboard handler
	// This is just for any additional logic if needed
}

func (a *App) handleRefresh() {
	if err := a.explorer.LoadFiles(); err != nil {
		return
	}
	
	a.allFiles = a.explorer.GetAllFiles()
	searchQuery := a.layoutManager.GetSearchComponent().GetText()
	a.filteredFiles = a.explorer.FilterFiles(searchQuery)
	
	a.updateAllComponents()
}

func (a *App) updateAllComponents() {
	// Update file list
	a.layoutManager.GetFileListComponent().Update(a.filteredFiles)
	
	// Update stats
	a.layoutManager.GetStatsComponent().Update(a.allFiles)
	a.layoutManager.GetStatsComponent().SetFilteredFiles(a.filteredFiles)
	
	// Update preview and details with first file if available
	if len(a.filteredFiles) > 0 {
		a.currentFile = &a.filteredFiles[0]
		a.layoutManager.GetPreviewComponent().Update(*a.currentFile)
		a.layoutManager.GetDetailsComponent().Update(*a.currentFile)
		a.layoutManager.GetStatusBarComponent().Update(*a.currentFile)
	} else {
		a.currentFile = nil
		a.layoutManager.GetDetailsComponent().SetNoFileSelected()
		a.layoutManager.GetStatusBarComponent().Update("")
	}
	
	// Update status bar counts
	a.layoutManager.GetStatusBarComponent().SetCounts(len(a.filteredFiles), len(a.allFiles))
}

func (a *App) Run() error {
	// Set initial focus
	a.keyHandler.SetCurrentFocus(a.config.InitialFocus)
	
	// Set application background to transparent and clear screen
	a.tvApp.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorDefault))
		screen.Clear()
		return false
	})
	
	return a.tvApp.SetRoot(a.layoutManager.GetRoot(), true).Run()
}