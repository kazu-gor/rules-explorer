package types

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type FileItem struct {
	Path    string
	Content string
}

type FileType int

const (
	CursorRule FileType = iota
	ClaudeConfig
	ConfigFile
	Unknown
)

func (ft FileType) String() string {
	switch ft {
	case CursorRule:
		return "Cursor Rule"
	case ClaudeConfig:
		return "Claude Config"
	case ConfigFile:
		return "Configuration"
	default:
		return "Unknown"
	}
}

type Focus int

const (
	FocusSearch Focus = iota
	FocusFileList
	FocusPreview
)

type EventType int

const (
	EventFileSelected EventType = iota
	EventFileChanged
	EventSearchChanged
	EventFocusChanged
	EventRefresh
	EventQuit
	EventEditFile
)

type Event struct {
	Type EventType
	Data interface{}
}

type FileEvent struct {
	File  FileItem
	Index int
}

type SearchEvent struct {
	Query string
}

type FocusEvent struct {
	Focus Focus
}

type EventHandler func(event Event)

type FileExplorer interface {
	LoadFiles() error
	FilterFiles(filter string) []FileItem
	GetAllFiles() []FileItem
}

type Component interface {
	GetPrimitive() tview.Primitive
	SetEventHandler(handler EventHandler)
	Focus()
	Blur()
	Update(data interface{})
}

type Theme interface {
	GetColors() ColorScheme
	GetIcons() IconSet
}

type ColorScheme struct {
	Primary     tcell.Color
	Secondary   tcell.Color
	Accent      tcell.Color
	Background  tcell.Color
	Text        tcell.Color
	Border      tcell.Color
	BorderFocus tcell.Color
	Success     tcell.Color
	Warning     tcell.Color
	Error       tcell.Color
}

type IconSet struct {
	CursorRule   string
	ClaudeConfig string
	ConfigFile   string
	Search       string
	File         string
	Folder       string
}