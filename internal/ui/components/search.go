package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"rules-explorer/internal/core/types"
)

type SearchComponent struct {
	input       *tview.InputField
	theme       types.Theme
	eventHandler types.EventHandler
}

func NewSearchComponent(th types.Theme) *SearchComponent {
	s := &SearchComponent{
		input: tview.NewInputField(),
		theme: th,
	}
	
	s.setupInput()
	return s
}

func (s *SearchComponent) setupInput() {
	colors := s.theme.GetColors()
	icons := s.theme.GetIcons()
	
	s.input.
		SetFieldWidth(0).
		SetPlaceholder("Type to filter files...").
		SetFieldBackgroundColor(tcell.ColorDefault).
		SetFieldTextColor(colors.Text).
		SetPlaceholderStyle(tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(colors.Secondary))
	
	s.input.SetBorder(true).
		SetTitle("[yellow]" + icons.Search + " Search Rules & Config Files[-]").
		SetTitleAlign(tview.AlignLeft).
		SetBorderColor(colors.Border).
		SetBackgroundColor(tcell.ColorDefault)
	
	s.input.SetChangedFunc(s.onSearchChanged)
}

func (s *SearchComponent) onSearchChanged(text string) {
	if s.eventHandler != nil {
		event := types.Event{
			Type: types.EventSearchChanged,
			Data: types.SearchEvent{Query: text},
		}
		s.eventHandler(event)
	}
}

func (s *SearchComponent) GetPrimitive() tview.Primitive {
	return s.input
}

func (s *SearchComponent) SetEventHandler(handler types.EventHandler) {
	s.eventHandler = handler
}

func (s *SearchComponent) Focus() {
	colors := s.theme.GetColors()
	s.input.SetBorderColor(colors.BorderFocus)
}

func (s *SearchComponent) Blur() {
	colors := s.theme.GetColors()
	s.input.SetBorderColor(colors.Border)
}

func (s *SearchComponent) Update(data interface{}) {
	// Search component doesn't need updates from external data
}

func (s *SearchComponent) GetText() string {
	return s.input.GetText()
}

func (s *SearchComponent) SetText(text string) {
	s.input.SetText(text)
}