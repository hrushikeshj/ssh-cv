package minesweeper

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Down):
			m.curPos.x = min(m.row-1, m.curPos.x+1)

		case key.Matches(msg, m.KeyMap.Up):
			m.curPos.x = max(0, m.curPos.x-1)

		case key.Matches(msg, m.KeyMap.Left):
			m.curPos.y = max(0, m.curPos.y-1)

		case key.Matches(msg, m.KeyMap.Right):
			m.curPos.y = min(m.col-1, m.curPos.y+1)

		case key.Matches(msg, m.KeyMap.Enter):
			if !m.hasWon() {
				m = m.unMask()
			}

		case key.Matches(msg, m.KeyMap.Flag):
			m = m.flag()

		case key.Matches(msg, m.KeyMap.Restart):
			m = NewGame(m.row, m.col, m.phy_width, m.r)
		}
	}
	return m, nil
}

func (m *Model) neighbors(p pos) []pos {
	nei := []pos{}

	for i := max(0, p.x-1); i <= min(m.row-1, p.x+1); i++ {
		for j := max(0, p.y-1); j <= min(m.col-1, p.y+1); j++ {
			if !(i == p.x && j == p.y) {
				nei = append(nei, pos{x: i, y: j})
			}
		}
	}

	return nei
}

func (m Model) unMask() Model {
	cell := &m.board[m.curPos.x][m.curPos.y]
	now_started := false

	if m.gameState == starting {
		// strat the game, by genreating mines
		m.generateMines()
		m.gameState = playing
		now_started = true
	}

	if cell.flag { // if flaged, just remove it
		return m.flag()
	}

	if cell.mine {
		m.gameState = out
		cell.shown = true
		return m
	}

	q := []pos{m.curPos}
	// first move
	if now_started {
		for _, neighbour := range m.neighbors(m.curPos) {
			cel := m.board[neighbour.x][neighbour.y]
			if !cel.shown && !cel.mine {
				q = append(q, neighbour)
			}
		}
	}

	for len(q) != 0 {
		var p pos
		p, q = q[0], q[1:]
		m.board[p.x][p.y].shown = true

		// unmask neighbours if mine_count is 0
		if m.board[p.x][p.y].mine_count == 0 {
			for _, neighbour := range m.neighbors(p) {
				cel := m.board[neighbour.x][neighbour.y]
				if !cel.shown && !cel.mine {
					q = append(q, neighbour)
				}
			}
		}
	}

	return m
}

func (m Model) flag() Model {
	cell := &m.board[m.curPos.x][m.curPos.y]

	idx := posToIdx(m.curPos, m.row, m.col)
	if cell.flag {
		delete(m.flags, idx)
		cell.flag = false
	} else {
		if len(m.flags) >= m.mine_count || cell.shown {
			// only `mine_count` flags can be planted
			// and flag can't be planted on unmasked cell
			return m
		}
		m.flags[idx] = 1
		cell.flag = true
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}
