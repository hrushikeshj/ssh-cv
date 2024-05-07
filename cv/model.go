package cv

import (
	"fmt"
	"os"
	"time"

	// "strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type Model struct {
	width          int
	physicalWidth  int
	physicalHeight int
	cvRendered     string
	styles         styles
	r              *lipgloss.Renderer
	viewport       viewport.Model
	spinner        spinner.Model
	loaded         bool
}

type readyMsg struct{}

var (
	p, h, _ = term.GetSize(int(os.Stdout.Fd()))
)

func (m *Model) updateWidthAndRender(phy_width, phy_height int) {
	m.physicalWidth = phy_width
	m.physicalHeight = phy_height
	m.width = min(MAX_WIDTH, m.physicalWidth)
	m.styles = m.makeStyle(m.r)
	m.width = min(MAX_WIDTH, m.physicalWidth)

	m.cvRendered = m.RenderCV()
	m.viewport.SetContent(m.cvRendered)
}

func NewModel() *Model {
	m := Model{
		width:         MAX_WIDTH,
		r:             lipgloss.DefaultRenderer(),
		cvRendered:    "",
		physicalWidth: p,
		loaded:        false,
	}

	vp := viewport.New(m.physicalWidth, 20)
	// vp.Style = lipgloss.NewStyle().
	// 	BorderStyle(lipgloss.RoundedBorder()).
	// 	BorderForeground(lipgloss.Color("62")).
	// 	PaddingRight(2)
	m.viewport = vp

	// spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	m.spinner = s

	// set width
	m.updateWidthAndRender(m.physicalWidth, h)
	return &m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(doTick(), m.spinner.Tick)
}

// loading screen for 600ms
func doTick() tea.Cmd {
	return tea.Tick(time.Millisecond*800, func(t time.Time) tea.Msg {
		return readyMsg{}
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	case tea.WindowSizeMsg:
		footerHeight := lipgloss.Height(m.footerView())
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - footerHeight
		m.updateWidthAndRender(msg.Width, msg.Height)
		return m, nil
	case readyMsg:
		m.loaded = true
		return m, nil
	case spinner.TickMsg:
		if m.loaded { // end animation, if loaded
			return m, nil
		}
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.MouseMsg:
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m Model) View() string {
	if m.loaded {
		return fmt.Sprintf("%s\n%s", m.viewport.View(), m.footerView())
	}

	return m.loadingScreen()
}

func (m Model) loadingScreen() string {
	h := (m.physicalHeight / 2) - 1
	var style = m.r.NewStyle().
		Bold(true).
		Width(m.physicalWidth).
		PaddingTop(h).
		Align(lipgloss.Center)
	return style.Render(m.spinner.View() + m.styles.loadingText(" cv.hrushi.dev"))
	// return m.r.NewStyle().
	// 			Width(50).
	// 			Height(60).
	// 			Align(lipgloss.Center).
	// 			Render(m.spinner.View() + " Loading!!")
}

var (
	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)
	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	statusText = lipgloss.NewStyle().Inherit(statusBarStyle).Foreground(lipgloss.Color("241"))

	fishCakeStyle = statusNugget.Copy().Background(special).Foreground(lipgloss.Color("#000000"))
)

func (m Model) footerView() string {
	w := lipgloss.Width

	statusKey := statusStyle.Render("RESUME")
	fishCake := fishCakeStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	statusVal := statusText.
		Width(m.physicalWidth - w(statusKey) - w(fishCake)).
		Render("↑/↓: Navigate • q: Quit")

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		fishCake,
	)

	return statusBarStyle.Width(m.physicalWidth).Render(bar)
}
