package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func HandleGlobalKeys(app *tview.Application, searchInput *tview.InputField, fileList *tview.List) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
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
	}
}

func HandleFileListKeys(fileList *tview.List) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
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
	}
}

func HandleSearchEscape(app *tview.Application) func(key tcell.Key) {
	return func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
	}
}