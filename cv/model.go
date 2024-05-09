package cv

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const LOADING_TIME = 800

type currentScreen int

const (
	cv currentScreen = iota
	links
	game
)

type Model struct {
	width            int
	physicalWidth    int
	physicalHeight   int
	cvRendered       string
	styles           styles
	r                *lipgloss.Renderer
	viewport         viewport.Model
	spinner          spinner.Model
	loaded           bool
	loadingScreenMsg string
	currentView      currentScreen
}

type halfLoadingTick struct{}
type readyMsg struct{}

func (m *Model) UpdateWidthAndRender(phy_width, phy_height int) {
	m.physicalWidth = phy_width
	m.physicalHeight = phy_height
	m.width = min(MAX_WIDTH, m.physicalWidth)
	m.styles = m.makeStyle(m.r)
	m.width = min(MAX_WIDTH, m.physicalWidth)

	m.cvRendered = m.RenderCV()
	m.viewport.SetContent(m.cvRendered)
}

func NewModel(width, height int, r *lipgloss.Renderer) *Model {
	m := Model{
		width:            MAX_WIDTH,
		r:                r,
		cvRendered:       "",
		physicalWidth:    width,
		physicalHeight:   height,
		loaded:           false,
		loadingScreenMsg: "cv.hrushi.dev",
		currentView:      cv,
	}

	vp := viewport.New(m.physicalWidth, 20)
	m.viewport = vp

	// spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = m.r.NewStyle().Foreground(lipgloss.Color("205"))
	m.spinner = s

	// set width
	m.UpdateWidthAndRender(width, height)
	return &m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(doTick(halfLoadingTick{}), m.spinner.Tick)
}

// loading screen for 600ms
func doTick(msg tea.Msg) tea.Cmd {
	return tea.Tick(time.Millisecond*LOADING_TIME/2, func(t time.Time) tea.Msg {
		return msg
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			switch m.currentView {
			case links:
				// go back to cv
				m.currentView = cv
				return m, nil
			}

			return m, tea.Quit
		case "l", "L":
			m.currentView = links
			return m, nil
		default:
			var cmd tea.Cmd
			m.viewport, cmd = m.viewport.Update(msg)
			return m, cmd
		}
	case tea.WindowSizeMsg:
		footerHeight := lipgloss.Height(m.footerView())
		m.viewport.Width = msg.Width
		m.viewport.Height = msg.Height - footerHeight
		m.UpdateWidthAndRender(msg.Width, msg.Height)
		return m, nil
	case halfLoadingTick:
		m.loadingScreenMsg = "Hello There!!"
		return m, doTick(readyMsg{})
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
	//return fmt.Sprintf("%s\n%s", m.linksView(), m.footerView())
	if m.loaded {
		var view string = "_"
		switch m.currentView {
		case cv:
			view = m.viewport.View()
		case links:
			view = m.linksView()
		}

		return fmt.Sprintf("%s\n%s", view, m.footerView())
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
	return style.Render(m.spinner.View() + m.styles.loadingText(m.loadingScreenMsg))
}

func (m Model) footerView() string {
	w := lipgloss.Width

	statusKey := m.styles.statusStyle.Render("RESUME")
	fishCake := m.styles.scrollPercent.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	statusVal := m.styles.statusText.
		Width(m.physicalWidth - w(statusKey) - w(fishCake)).
		Render("↑/↓: Navigate • l: Links • esc: Back • q: Quit")

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
		fishCake,
	)

	return m.styles.statusBarStyle.Width(m.physicalWidth).Render(bar)
}
