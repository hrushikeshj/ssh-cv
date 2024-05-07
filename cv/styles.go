package cv

import (
	"github.com/charmbracelet/lipgloss"
	// "golang.org/x/term"
	"math"
	// "os"
)

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

type styles struct {
	body             lipgloss.Style
	nameTitle        lipgloss.Style
	sectionHeader    func(strs ...string) string
	subSectionHeader lipgloss.Style
	duration         lipgloss.Style
	companyDisplay   lipgloss.Style
	sectionBlock     func(strs ...string) string
	bulletBlock      lipgloss.Style
	techText         lipgloss.Style
	bulletWidth      lipgloss.Style
	loadingText      func(strs ...string) string
}

var (
	// light: #878787
	// #FFFFFF
	// brown #FF5F00

	subtle           = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight        = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special          = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	loadingTextColor = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	lightText        = lipgloss.Color("#878787")
)

const MAX_WIDTH = 80 //75

func (m *Model) makeStyle(r *lipgloss.Renderer) styles {
	//physicalWidth, _, _ = term.GetSize(int(os.Stdout.Fd()))

	_width := m.width
	padding := int(math.Floor((float64(m.physicalWidth) - float64(_width)) / 2))
	return styles{
		body: r.NewStyle().Width(_width).
			MarginLeft(padding).
			MarginRight(padding).
			Width(_width),
		nameTitle: r.NewStyle().
			Foreground(special).
			Width(_width).
			Align(lipgloss.Center),
		sectionHeader: r.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(special).
			Width(_width).
			Render,
		subSectionHeader: r.NewStyle().
			Bold(true),
		duration:       r.NewStyle().Faint(true),
		companyDisplay: r.NewStyle().Italic(true),
		sectionBlock:   r.NewStyle().MarginLeft(2).Render,
		bulletBlock:    r.NewStyle().MarginLeft(1),
		techText:       r.NewStyle().Faint(true).Italic(true),
		bulletWidth:    r.NewStyle().Width(_width - 5).Foreground(lightText),
		loadingText:    r.NewStyle().Foreground(loadingTextColor).Render,
	}
}
