package components

import (
	"fmt"
	"time"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"rules-explorer/internal/core/types"
	"rules-explorer/internal/ui/theme"
)

type StatsComponent struct {
	textView     *tview.TextView
	theme        types.Theme
	eventHandler types.EventHandler
	allFiles     []types.FileItem
	filteredFiles []types.FileItem
}

func NewStatsComponent(th types.Theme) *StatsComponent {
	s := &StatsComponent{
		textView: tview.NewTextView(),
		theme:    th,
	}
	
	s.setupTextView()
	return s
}

func (s *StatsComponent) setupTextView() {
	colors := s.theme.GetColors()
	
	// Create transparent text style
	transparentStyle := tcell.StyleDefault.
		Background(tcell.ColorDefault).
		Foreground(colors.Text)
	
	s.textView.
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetTextStyle(transparentStyle)
	
	s.textView.SetBorder(true).
		SetTitle("[magenta]📊 Statistics[-]").
		SetTitleAlign(tview.AlignLeft).
		SetBorderColor(colors.Border).
		SetBackgroundColor(tcell.ColorDefault)
}

func (s *StatsComponent) GetPrimitive() tview.Primitive {
	return s.textView
}

func (s *StatsComponent) SetEventHandler(handler types.EventHandler) {
	s.eventHandler = handler
}

func (s *StatsComponent) Focus() {
	colors := s.theme.GetColors()
	s.textView.SetBorderColor(colors.BorderFocus)
}

func (s *StatsComponent) Blur() {
	colors := s.theme.GetColors()
	s.textView.SetBorderColor(colors.Border)
}

func (s *StatsComponent) Update(data interface{}) {
	switch v := data.(type) {
	case []types.FileItem:
		s.allFiles = v
		s.updateStats()
	}
}

func (s *StatsComponent) SetFilteredFiles(files []types.FileItem) {
	s.filteredFiles = files
	s.updateStats()
}

func (s *StatsComponent) updateStats() {
	totalCount := len(s.allFiles)
	filteredCount := len(s.filteredFiles)
	
	// Count by type
	cursorRules := 0
	claudeConfigs := 0
	configFiles := 0
	
	for _, file := range s.allFiles {
		fileType := theme.DetermineFileType(file.Path)
		switch fileType {
		case types.CursorRule:
			cursorRules++
		case types.ClaudeConfig:
			claudeConfigs++
		case types.ConfigFile:
			configFiles++
		}
	}
	
	icons := s.theme.GetIcons()
	
	stats := fmt.Sprintf(`[yellow]Total Files:[-] %d
[yellow]Filtered:[-] %d

[yellow]By Type:[-]
[red]%s[-] Cursor Rules: %d
[green]%s[-] Claude Configs: %d
[blue]%s[-] Config Files: %d

[yellow]Timestamp:[-]
%s`,
		totalCount,
		filteredCount,
		icons.CursorRule, cursorRules,
		icons.ClaudeConfig, claudeConfigs,
		icons.ConfigFile, configFiles,
		time.Now().Format("15:04:05"))
	
	s.textView.SetText(stats)
}