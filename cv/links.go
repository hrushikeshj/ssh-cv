package cv

import "github.com/charmbracelet/lipgloss"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxLen(links []link) int {
	m := 0
	for _, link := range links {
		m = max(m, len(link.name))
	}

	return m
}

// mountains moving below
func (m Model) linksView() string {
	var linksR []string
	mx := maxLen(MyCV.links)
	for _, link := range MyCV.links {
		name := m.r.NewStyle().Width(mx).Foreground(lightText).Render(link.name)
		div := m.r.NewStyle().Foreground(lightText).Render(" : ")

		linksR = append(
			linksR,
			lipgloss.JoinHorizontal(lipgloss.Top, name+div, m.styles.urlStyle(link.url)),
		)
	}

	s := lipgloss.JoinVertical(lipgloss.Left, linksR...)

	return lipgloss.Place(m.physicalWidth, m.physicalHeight-lipgloss.Height(m.footerView()),
		lipgloss.Center, lipgloss.Center,
		m.styles.linksBox.Render(s),
		// lipgloss.WithWhitespaceChars("猫咪"),
		// lipgloss.WithWhitespaceForeground(subtle),
	)
}
