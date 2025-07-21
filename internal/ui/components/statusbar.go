package components

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"rules-explorer/internal/core/types"
	"rules-explorer/internal/utils"
)

type StatusBarComponent struct {
	textView      *tview.TextView
	theme         types.Theme
	eventHandler  types.EventHandler
	currentFile   string
	filteredCount int
	totalCount    int
}

func NewStatusBarComponent(th types.Theme) *StatusBarComponent {
	s := &StatusBarComponent{
		textView:    tview.NewTextView(),
		theme:       th,
		currentFile: "None",
	}
	
	s.setupTextView()
	return s
}

func (s *StatusBarComponent) setupTextView() {
	colors := s.theme.GetColors()
	
	// Create transparent text style
	transparentStyle := tcell.StyleDefault.
		Background(tcell.ColorDefault).
		Foreground(colors.Text)
	
	s.textView.
		SetDynamicColors(true).
		SetRegions(false).
		SetWordWrap(false).
		SetTextStyle(transparentStyle)
	
	s.textView.SetBackgroundColor(tcell.ColorDefault)
	s.updateStatus()
}

func (s *StatusBarComponent) GetPrimitive() tview.Primitive {
	return s.textView
}

func (s *StatusBarComponent) SetEventHandler(handler types.EventHandler) {
	s.eventHandler = handler
}

func (s *StatusBarComponent) Focus() {
	// Status bar doesn't get focus
}

func (s *StatusBarComponent) Blur() {
	// Status bar doesn't get focus
}

func (s *StatusBarComponent) Update(data interface{}) {
	switch v := data.(type) {
	case types.FileItem:
		s.currentFile = utils.GetBaseName(v.Path)
		s.updateStatus()
	case string:
		if v == "" {
			s.currentFile = "None"
		} else {
			s.currentFile = v
		}
		s.updateStatus()
	}
}

func (s *StatusBarComponent) SetCounts(filtered, total int) {
	s.filteredCount = filtered
	s.totalCount = total
	s.updateStatus()
}

func (s *StatusBarComponent) updateStatus() {
	statusText := fmt.Sprintf(" [yellow]Rules Explorer[-] | Files: %d/%d | Current: [aqua]%s[-] | [white]Tab[-]: Switch Panes | [white]Esc[-]: Exit",
		s.filteredCount, s.totalCount, s.currentFile)
	
	s.textView.SetText(statusText)
}