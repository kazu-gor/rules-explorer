package input

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"rules-explorer/internal/core/types"
	"rules-explorer/internal/ui/components"
)

type KeyboardHandler struct {
	app          *tview.Application
	eventHandler types.EventHandler
	currentFocus types.Focus
	components   map[types.Focus]types.Component
}

func NewKeyboardHandler(app *tview.Application) *KeyboardHandler {
	return &KeyboardHandler{
		app:          app,
		currentFocus: types.FocusSearch,
		components:   make(map[types.Focus]types.Component),
	}
}

func (k *KeyboardHandler) SetEventHandler(handler types.EventHandler) {
	k.eventHandler = handler
}

func (k *KeyboardHandler) RegisterComponent(focus types.Focus, component types.Component) {
	k.components[focus] = component
}

func (k *KeyboardHandler) HandleGlobalKeys(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyTab:
		k.switchFocus(true)
		return nil
	case tcell.KeyBacktab:
		k.switchFocus(false)
		return nil
	case tcell.KeyRune:
		switch event.Rune() {
		case 'l':
			k.switchFocus(true)
			return nil
		case 'h':
			k.switchFocus(false)
			return nil
		case 'q':
			if k.eventHandler != nil {
				k.eventHandler(types.Event{
					Type: types.EventQuit,
					Data: nil,
				})
			}
			k.app.Stop()
			return nil
		}
	case tcell.KeyCtrlC, tcell.KeyEscape:
		if k.eventHandler != nil {
			k.eventHandler(types.Event{
				Type: types.EventQuit,
				Data: nil,
			})
		}
		k.app.Stop()
		return nil
	case tcell.KeyEnter:
		k.handleEnter()
		return nil
	}
	
	// Handle navigation in file list
	if k.currentFocus == types.FocusFileList {
		switch event.Key() {
		case tcell.KeyDown, tcell.KeyCtrlN:
			k.handleFileListNavigation(1)
			return nil
		case tcell.KeyUp, tcell.KeyCtrlP:
			k.handleFileListNavigation(-1)
			return nil
		case tcell.KeyRune:
			switch event.Rune() {
			case 'j':
				k.handleFileListNavigation(1)
				return nil
			case 'k':
				k.handleFileListNavigation(-1)
				return nil
			case 'e':
				if k.eventHandler != nil {
					k.eventHandler(types.Event{
						Type: types.EventEditFile,
						Data: nil,
					})
				}
				return nil
			}
		}
	}
	
	return event
}

func (k *KeyboardHandler) switchFocus(forward bool) {
	// Blur current component
	if comp, exists := k.components[k.currentFocus]; exists {
		comp.Blur()
	}
	
	// Switch focus
	if forward {
		switch k.currentFocus {
		case types.FocusSearch:
			k.currentFocus = types.FocusFileList
		case types.FocusFileList:
			k.currentFocus = types.FocusSearch
		case types.FocusPreview:
			k.currentFocus = types.FocusSearch
		}
	} else {
		// Backward focus (h key / Shift+Tab)
		switch k.currentFocus {
		case types.FocusSearch:
			k.currentFocus = types.FocusFileList
		case types.FocusFileList:
			k.currentFocus = types.FocusSearch
		case types.FocusPreview:
			k.currentFocus = types.FocusFileList
		}
	}
	
	// Focus new component
	if comp, exists := k.components[k.currentFocus]; exists {
		comp.Focus()
		k.app.SetFocus(comp.GetPrimitive())
	}
	
	// Notify about focus change
	if k.eventHandler != nil {
		event := types.Event{
			Type: types.EventFocusChanged,
			Data: types.FocusEvent{Focus: k.currentFocus},
		}
		k.eventHandler(event)
	}
}

func (k *KeyboardHandler) handleEnter() {
	if k.currentFocus == types.FocusFileList {
		// File selection will be handled by the file list component itself
		// through its SetSelectedFunc callback
	}
}

func (k *KeyboardHandler) handleFileListNavigation(direction int) {
	if comp, exists := k.components[types.FocusFileList]; exists {
		if fileList, ok := comp.(*components.FileListComponent); ok {
			if direction > 0 {
				fileList.NavigateDown()
			} else {
				fileList.NavigateUp()
			}
		}
	}
}

func (k *KeyboardHandler) SetCurrentFocus(focus types.Focus) {
	k.currentFocus = focus
	
	// Blur all components first
	for _, comp := range k.components {
		comp.Blur()
	}
	
	// Focus the selected component
	if comp, exists := k.components[focus]; exists {
		comp.Focus()
		k.app.SetFocus(comp.GetPrimitive())
	}
}