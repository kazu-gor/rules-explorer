package components

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"rules-explorer/internal/core/types"
	"rules-explorer/internal/ui/theme"
	"rules-explorer/internal/utils"
)

type DetailsComponent struct {
	textView     *tview.TextView
	theme        types.Theme
	eventHandler types.EventHandler
}

func NewDetailsComponent(th types.Theme) *DetailsComponent {
	d := &DetailsComponent{
		textView: tview.NewTextView(),
		theme:    th,
	}
	
	d.setupTextView()
	return d
}

func (d *DetailsComponent) setupTextView() {
	colors := d.theme.GetColors()
	
	// Create transparent text style
	transparentStyle := tcell.StyleDefault.
		Background(tcell.ColorDefault).
		Foreground(colors.Text)
	
	d.textView.
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetScrollable(true).
		SetTextStyle(transparentStyle)
	
	d.textView.SetBorder(true).
		SetTitle("[green]📄 File Details[-]").
		SetTitleAlign(tview.AlignLeft).
		SetBorderColor(colors.Border).
		SetBackgroundColor(tcell.ColorDefault)
}

func (d *DetailsComponent) GetPrimitive() tview.Primitive {
	return d.textView
}

func (d *DetailsComponent) SetEventHandler(handler types.EventHandler) {
	d.eventHandler = handler
}

func (d *DetailsComponent) Focus() {
	colors := d.theme.GetColors()
	d.textView.SetBorderColor(colors.BorderFocus)
}

func (d *DetailsComponent) Blur() {
	colors := d.theme.GetColors()
	d.textView.SetBorderColor(colors.Border)
}

func (d *DetailsComponent) Update(data interface{}) {
	if file, ok := data.(types.FileItem); ok {
		d.updateFileDetails(file)
	}
}

func (d *DetailsComponent) updateFileDetails(file types.FileItem) {
	size := len(file.Content)
	lineCount := utils.CountLines(file.Content)
	
	sizeStr := utils.FormatFileSize(size)
	fileType := theme.DetermineFileType(file.Path)
	icons := d.theme.GetIcons()
	icon := theme.GetFileTypeIcon(fileType, icons)
	
	details := fmt.Sprintf(`%s [white]%s[-]

[yellow]Path:[-] %s
[yellow]Type:[-] %s
[yellow]Size:[-] %s
[yellow]Lines:[-] %d

[yellow]Content Preview:[-]
[gray]%s[-]`,
		icon, utils.GetBaseName(file.Path),
		file.Path,
		fileType.String(),
		sizeStr,
		lineCount,
		utils.GetContentPreview(file.Content, 5, 50))
	
	d.textView.SetText(details)
}

func (d *DetailsComponent) SetNoFileSelected() {
	d.textView.SetText("[yellow]No files selected[-]")
}