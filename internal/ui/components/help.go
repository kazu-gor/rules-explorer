package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"rules-explorer/internal/core/types"
)

type HelpComponent struct {
	textView     *tview.TextView
	theme        types.Theme
	eventHandler types.EventHandler
}

func NewHelpComponent(th types.Theme) *HelpComponent {
	h := &HelpComponent{
		textView: tview.NewTextView(),
		theme:    th,
	}
	
	h.setupTextView()
	return h
}

func (h *HelpComponent) setupTextView() {
	colors := h.theme.GetColors()
	icons := h.theme.GetIcons()
	
	// Create transparent text style
	transparentStyle := tcell.StyleDefault.
		Background(tcell.ColorDefault).
		Foreground(colors.Text)
	
	h.textView.
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetTextStyle(transparentStyle)
	
	h.textView.SetBorder(true).
		SetTitle("[blue]❓ Help & Commands[-]").
		SetTitleAlign(tview.AlignLeft).
		SetBorderColor(colors.Border).
		SetBackgroundColor(tcell.ColorDefault)
	
	helpText := `[yellow]Navigation:[-]
[white]Tab/l[-]     - Next pane
[white]Shift+Tab/h[-] - Previous pane
[white]↑/↓/j/k[-]   - Navigate files
[white]Ctrl+P/N[-]  - Navigate files
[white]Enter[-]     - Select file
[white]e[-]         - Edit file
[white]q/Esc[-]     - Exit
[white]Ctrl+C[-]    - Quit

[yellow]File Types:[-]
[red]` + icons.CursorRule + `[-] Cursor Rules (.mdc)
[green]` + icons.ClaudeConfig + `[-] Claude Config (CLAUDE.md)
[blue]` + icons.ConfigFile + `[-]  Config (.claude/*)`
	
	h.textView.SetText(helpText)
}

func (h *HelpComponent) GetPrimitive() tview.Primitive {
	return h.textView
}

func (h *HelpComponent) SetEventHandler(handler types.EventHandler) {
	h.eventHandler = handler
}

func (h *HelpComponent) Focus() {
	colors := h.theme.GetColors()
	h.textView.SetBorderColor(colors.BorderFocus)
}

func (h *HelpComponent) Blur() {
	colors := h.theme.GetColors()
	h.textView.SetBorderColor(colors.Border)
}

func (h *HelpComponent) Update(data interface{}) {
	// Help component doesn't need updates from external data
}