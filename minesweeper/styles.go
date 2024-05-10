package minesweeper

import (
	"github.com/charmbracelet/lipgloss"
	// "math"
)

type styles struct {
	flagStyle func(...string) string
	winBorder lipgloss.Style
	lostBorder lipgloss.Style
	focusedStyle lipgloss.Style
	hiddenCell lipgloss.Style
}

var (
	highlight        = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special          = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
)

func (m *Model) makeStyle(r *lipgloss.Renderer) styles {
	return styles{
		flagStyle: r.NewStyle().Foreground(lipgloss.Color("#FF5F87")).Render,
		winBorder: r.NewStyle().Foreground(special),
		lostBorder: r.NewStyle().Foreground(lipgloss.Color("#a85151")),
		focusedStyle: r.NewStyle().Background(highlight),
		hiddenCell: r.NewStyle().Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#b8b4b4"}),
	}
}
