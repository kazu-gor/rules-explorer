package components

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"rules-explorer/internal/core/types"
	"rules-explorer/internal/ui/theme"
	"rules-explorer/internal/utils"
)

type FileListComponent struct {
	list         *tview.List
	theme        types.Theme
	eventHandler types.EventHandler
	files        []types.FileItem
}

func NewFileListComponent(th types.Theme) *FileListComponent {
	f := &FileListComponent{
		list:  tview.NewList(),
		theme: th,
		files: make([]types.FileItem, 0),
	}
	
	f.setupList()
	return f
}

func (f *FileListComponent) setupList() {
	colors := f.theme.GetColors()
	icons := f.theme.GetIcons()
	
	// Create transparent styles
	selectedStyle := tcell.StyleDefault.
		Background(tcell.ColorDefault).
		Foreground(tcell.ColorYellow)
	
	mainTextStyle := tcell.StyleDefault.
		Background(tcell.ColorDefault).
		Foreground(colors.Text)
	
	secondaryTextStyle := tcell.StyleDefault.
		Background(tcell.ColorDefault).
		Foreground(colors.Secondary)
	
	f.list.
		ShowSecondaryText(true).
		SetSelectedFunc(f.onFileSelected).
		SetChangedFunc(f.onFileChanged).
		SetMainTextColor(colors.Text).
		SetSecondaryTextColor(colors.Secondary).
		SetMainTextStyle(mainTextStyle).
		SetSecondaryTextStyle(secondaryTextStyle).
		SetSelectedBackgroundColor(tcell.ColorDefault).
		SetSelectedTextColor(tcell.ColorYellow).
		SetSelectedStyle(selectedStyle).
		SetSelectedFocusOnly(false).
		SetHighlightFullLine(true).
		SetUseStyleTags(false, false)
	
	f.list.SetBorder(true).
		SetTitle("[aqua]" + icons.Folder + " Files[-]").
		SetTitleAlign(tview.AlignLeft).
		SetBorderColor(colors.Border).
		SetBackgroundColor(tcell.ColorDefault)
}

func (f *FileListComponent) onFileSelected(index int, mainText string, secondaryText string, shortcut rune) {
	if f.eventHandler != nil && index < len(f.files) {
		event := types.Event{
			Type: types.EventFileSelected,
			Data: types.FileEvent{
				File:  f.files[index],
				Index: index,
			},
		}
		f.eventHandler(event)
	}
}

func (f *FileListComponent) onFileChanged(index int, mainText string, secondaryText string, shortcut rune) {
	if f.eventHandler != nil && index >= 0 && index < len(f.files) {
		event := types.Event{
			Type: types.EventFileChanged,
			Data: types.FileEvent{
				File:  f.files[index],
				Index: index,
			},
		}
		f.eventHandler(event)
	}
}

func (f *FileListComponent) GetPrimitive() tview.Primitive {
	return f.list
}

func (f *FileListComponent) SetEventHandler(handler types.EventHandler) {
	f.eventHandler = handler
}

func (f *FileListComponent) Focus() {
	colors := f.theme.GetColors()
	f.list.SetBorderColor(colors.BorderFocus)
	
	// Update selected style when focused
	selectedStyle := tcell.StyleDefault.
		Background(tcell.ColorDefault).
		Foreground(tcell.ColorYellow).
		Bold(true)
	f.list.SetSelectedStyle(selectedStyle)
}

func (f *FileListComponent) Blur() {
	colors := f.theme.GetColors()
	f.list.SetBorderColor(colors.Border)
	
	// Update selected style when blurred
	selectedStyle := tcell.StyleDefault.
		Background(tcell.ColorDefault).
		Foreground(colors.Secondary)
	f.list.SetSelectedStyle(selectedStyle)
}

func (f *FileListComponent) Update(data interface{}) {
	if files, ok := data.([]types.FileItem); ok {
		f.updateFiles(files)
	}
}

func (f *FileListComponent) updateFiles(files []types.FileItem) {
	f.files = files
	f.list.Clear()
	
	icons := f.theme.GetIcons()
	
	for _, file := range files {
		fileTypeEnum := theme.DetermineFileType(file.Path)
		icon := theme.GetFileTypeIconPlain(fileTypeEnum, icons)
		shortPath := utils.GetShortPath(file.Path, 40)
		fileType := fileTypeEnum.String()
		
		f.list.AddItem(
			fmt.Sprintf("%s %s", icon, shortPath),
			fileType, // Remove color tags completely
			0,
			nil,
		)
	}
	
	if len(files) > 0 {
		f.list.SetCurrentItem(0)
	}
}

func (f *FileListComponent) GetCurrentItem() int {
	return f.list.GetCurrentItem()
}

func (f *FileListComponent) SetCurrentItem(index int) {
	f.list.SetCurrentItem(index)
}

func (f *FileListComponent) GetItemCount() int {
	return f.list.GetItemCount()
}

func (f *FileListComponent) NavigateUp() {
	current := f.list.GetCurrentItem()
	if current > 0 {
		f.list.SetCurrentItem(current - 1)
	}
}

func (f *FileListComponent) NavigateDown() {
	current := f.list.GetCurrentItem()
	if current < f.list.GetItemCount()-1 {
		f.list.SetCurrentItem(current + 1)
	}
}