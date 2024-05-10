package minesweeper

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func mineCountColor(i int, r *lipgloss.Renderer) string {
	var s string
	colors := map[int]string{
		1: "#2196f3",
		2: "#FF5F87",
		3: "#f40000",
		4: "#7b00d2",
		5: "#10becd",
		6: "#d51880",
		7: "#ffc107",
		8: "#ff5722",
	}

	color, ok := colors[i]
	if ok {
		s = r.NewStyle().
			Foreground(lipgloss.Color(color)).
			Render(fmt.Sprintf("%d", i))
	} else {
		s = " "
	}

	return s
}

func (m Model) BoardTable() string {
	board := make([][]string, m.row)
	for i := 0; i < m.row; i++ {
		board[i] = make([]string, m.col)
	}


	for i := 0; i < m.row; i++ {
		for j := 0; j < m.col; j++ {
			cell := m.board[i][j]
			s := " "

			style := m.r.NewStyle()
			if m.curPos.x == i && m.curPos.y == j {
				style = m.styles.focusedStyle
			} else if !cell.shown && !cell.flag {
				style = m.styles.hiddenCell
			}

			if cell.shown {
				s = mineCountColor(cell.mine_count, m.r)
			}
			if cell.flag {
				s = m.styles.flagStyle("⚑")
			}

			if cell.mine && cell.shown {
				s = "֎"
			}

			board[i][j] = style.Render(s)
		}
	}

	borderStyle := m.r.NewStyle()
	if m.hasWon() {
		borderStyle = m.styles.winBorder
	} else if m.gameState == out {
		borderStyle = m.styles.lostBorder
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(borderStyle).
		BorderRow(true).
		BorderColumn(true).
		Rows(board...).
		StyleFunc(func(row, col int) lipgloss.Style {
			return m.r.NewStyle().Padding(0, 1)
		})

	return t.Render()
}

func (m Model) View() string {
	center := m.r.NewStyle().
		Width(m.phy_width).
		Align(lipgloss.Center).Render
	msg := ""

	if m.hasWon() {
		msg = " You won!! Press [r] to play again."
	} else if m.gameState == out {
		msg = "You lost :( T[r]y again"
	}

	top := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.styles.flagStyle("⚑"),
		fmt.Sprintf(" %d", m.mine_count-len(m.flags)),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		center(top),
		center(msg),
		center(m.BoardTable()),
		"",
		center(m.help.View(m.KeyMap)),
	)
}

func (m Model) BoardTableDebug() string {
	board := make([][]string, m.row)
	for i := 0; i < m.row; i++ {
		board[i] = make([]string, m.col)
	}

	for i := 0; i < m.row; i++ {
		for j := 0; j < m.col; j++ {
			cell := m.board[i][j]
			s := ""
			if m.curPos.x == i && m.curPos.y == j {
				s += "* "
			}
			if cell.shown {
				s += "• "
			}
			if cell.flag {
				s += "⚑ "
			}
			if m.board[i][j].mine {
				board[i][j] = fmt.Sprintf("M %d %s", cell.mine_count, s)
			} else {
				board[i][j] = fmt.Sprintf("%d %s", cell.mine_count, s)
			}
		}
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		BorderColumn(true).
		Rows(board...).
		StyleFunc(func(row, col int) lipgloss.Style {
			return m.r.NewStyle().Padding(0, 1)
		})

	// `lipgloss.Left`. edge will be inline
	s := fmt.Sprintf("State: %d, won: %t, flags: %d[%d], x: %d, y: %d ffffffff ffffffffff f", m.gameState,
		m.hasWon(), len(m.flags), m.mine_count-len(m.flags), m.curPos.x, m.curPos.y)
	return lipgloss.JoinVertical(lipgloss.Left, t.Render(), s)
}
