package theme

import (
	"github.com/gdamore/tcell/v2"
	"rules-explorer/internal/core/types"
)

type DefaultTheme struct{}

func New() types.Theme {
	return &DefaultTheme{}
}

func (t *DefaultTheme) GetColors() types.ColorScheme {
	return types.ColorScheme{
		Primary:     tcell.ColorAqua,
		Secondary:   tcell.ColorGray,
		Accent:      tcell.ColorYellow,
		Background:  tcell.ColorDefault,
		Text:        tcell.ColorWhite,
		Border:      tcell.NewHexColor(0xC8D3F5),
		BorderFocus: tcell.ColorTeal,
		Success:     tcell.ColorGreen,
		Warning:     tcell.ColorYellow,
		Error:       tcell.ColorRed,
	}
}

func (t *DefaultTheme) GetIcons() types.IconSet {
	return types.IconSet{
		CursorRule:   "📋",
		ClaudeConfig: "📝",
		ConfigFile:   "📒",
		Search:       "🔍",
		File:         "📄",
		Folder:       "📁",
	}
}

func GetFileTypeColor(fileType types.FileType) string {
	switch fileType {
	case types.CursorRule:
		return "[red]"
	case types.ClaudeConfig:
		return "[green]"
	case types.ConfigFile:
		return "[blue]"
	default:
		return "[white]"
	}
}

func GetFileTypeIcon(fileType types.FileType, icons types.IconSet) string {
	color := GetFileTypeColor(fileType)
	switch fileType {
	case types.CursorRule:
		return color + ":-]" + icons.CursorRule + "[:-]"
	case types.ClaudeConfig:
		return color + ":-]" + icons.ClaudeConfig + "[:-]"
	case types.ConfigFile:
		return color + ":-]" + icons.ConfigFile + "[:-]"
	default:
		return color + ":-]" + icons.File + "[:-]"
	}
}

func GetFileTypeIconPlain(fileType types.FileType, icons types.IconSet) string {
	switch fileType {
	case types.CursorRule:
		return icons.CursorRule
	case types.ClaudeConfig:
		return icons.ClaudeConfig
	case types.ConfigFile:
		return icons.ConfigFile
	default:
		return icons.File
	}
}
