package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"rules-explorer/internal/core/types"
)

type PreviewComponent struct {
	textView     *tview.TextView
	theme        types.Theme
	eventHandler types.EventHandler
}

func NewPreviewComponent(th types.Theme) *PreviewComponent {
	p := &PreviewComponent{
		textView: tview.NewTextView(),
		theme:    th,
	}
	
	p.setupTextView()
	return p
}

func (p *PreviewComponent) setupTextView() {
	colors := p.theme.GetColors()
	
	// Create transparent text style
	transparentStyle := tcell.StyleDefault.
		Background(tcell.ColorDefault).
		Foreground(colors.Text)
	
	p.textView.
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetScrollable(true).
		SetTextStyle(transparentStyle)
	
	p.textView.SetBorder(true).
		SetTitle("[white]📖 Content Preview[-]").
		SetTitleAlign(tview.AlignLeft).
		SetBorderColor(colors.Border).
		SetBackgroundColor(tcell.ColorDefault)
}

func (p *PreviewComponent) GetPrimitive() tview.Primitive {
	return p.textView
}

func (p *PreviewComponent) SetEventHandler(handler types.EventHandler) {
	p.eventHandler = handler
}

func (p *PreviewComponent) Focus() {
	colors := p.theme.GetColors()
	p.textView.SetBorderColor(colors.BorderFocus)
}

func (p *PreviewComponent) Blur() {
	colors := p.theme.GetColors()
	p.textView.SetBorderColor(colors.Border)
}

func (p *PreviewComponent) Update(data interface{}) {
	switch v := data.(type) {
	case types.FileItem:
		p.SetContent(v.Content)
	case string:
		p.SetContent(v)
	}
}

func (p *PreviewComponent) SetContent(content string) {
	p.textView.Clear()
	p.textView.SetText(content)
}

func (p *PreviewComponent) Clear() {
	p.textView.Clear()
}