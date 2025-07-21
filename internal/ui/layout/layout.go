package layout

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"rules-explorer/internal/core/types"
	"rules-explorer/internal/ui/components"
)

type Manager struct {
	theme      types.Theme
	root       *tview.Flex
	mainLayout *tview.Flex
	
	// Components
	search    *components.SearchComponent
	fileList  *components.FileListComponent
	preview   *components.PreviewComponent
	details   *components.DetailsComponent
	stats     *components.StatsComponent
	help      *components.HelpComponent
	statusBar *components.StatusBarComponent
}

func NewManager(theme types.Theme) *Manager {
	m := &Manager{
		theme: theme,
	}
	
	m.createComponents()
	m.setupLayout()
	
	return m
}

func (m *Manager) createComponents() {
	m.search = components.NewSearchComponent(m.theme)
	m.fileList = components.NewFileListComponent(m.theme)
	m.preview = components.NewPreviewComponent(m.theme)
	m.details = components.NewDetailsComponent(m.theme)
	m.stats = components.NewStatsComponent(m.theme)
	m.help = components.NewHelpComponent(m.theme)
	m.statusBar = components.NewStatusBarComponent(m.theme)
}

func (m *Manager) setupLayout() {
	// Left side: Search + File List
	leftPanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(m.search.GetPrimitive(), 3, 0, false).
		AddItem(m.fileList.GetPrimitive(), 0, 1, false)
	leftPanel.SetBackgroundColor(tcell.ColorDefault)
	
	// Main content area: Left (Files) | Right (Preview)
	mainContent := tview.NewFlex().
		AddItem(leftPanel, 0, 1, false).
		AddItem(m.preview.GetPrimitive(), 0, 2, false)
	mainContent.SetBackgroundColor(tcell.ColorDefault)
	
	// Bottom info panel: File Details | Stats | Help
	bottomInfo := tview.NewFlex().
		AddItem(m.details.GetPrimitive(), 0, 1, false).
		AddItem(m.stats.GetPrimitive(), 0, 1, false).
		AddItem(m.help.GetPrimitive(), 0, 1, false)
	bottomInfo.SetBackgroundColor(tcell.ColorDefault)
	
	// Main layout: Top (Files + Preview) | Bottom (Info panels)
	m.mainLayout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mainContent, 0, 3, false).
		AddItem(bottomInfo, 0, 1, false)
	m.mainLayout.SetBackgroundColor(tcell.ColorDefault)
	
	// Overall layout with status bar
	m.root = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(m.mainLayout, 0, 1, true).
		AddItem(m.statusBar.GetPrimitive(), 1, 0, false)
	m.root.SetBackgroundColor(tcell.ColorDefault)
}

func (m *Manager) GetRoot() tview.Primitive {
	return m.root
}

func (m *Manager) GetComponents() map[types.Focus]types.Component {
	return map[types.Focus]types.Component{
		types.FocusSearch:   m.search,
		types.FocusFileList: m.fileList,
		types.FocusPreview:  m.preview,
	}
}

// Component accessors
func (m *Manager) GetSearchComponent() *components.SearchComponent {
	return m.search
}

func (m *Manager) GetFileListComponent() *components.FileListComponent {
	return m.fileList
}

func (m *Manager) GetPreviewComponent() *components.PreviewComponent {
	return m.preview
}

func (m *Manager) GetDetailsComponent() *components.DetailsComponent {
	return m.details
}

func (m *Manager) GetStatsComponent() *components.StatsComponent {
	return m.stats
}

func (m *Manager) GetHelpComponent() *components.HelpComponent {
	return m.help
}

func (m *Manager) GetStatusBarComponent() *components.StatusBarComponent {
	return m.statusBar
}