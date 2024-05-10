package minesweeper

import (
	"log"
	"math/rand/v2"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

type gameState int

const MINE_PERCENTAGE = 5
const (
	starting gameState = iota
	playing
	out
)

type cell struct {
	mine       bool
	shown      bool
	mine_count int
	flag       bool
}

type board [][]cell
type pos struct {
	x int
	y int
}

type Model struct {
	row        int
	col        int
	r          *lipgloss.Renderer
	mine_count int
	board      board
	curPos     pos
	KeyMap     keyMap
	gameState  gameState
	flags      map[int]interface{}
	phy_width  int
	styles     styles
	help       help.Model
}

func NewGame(row, col, phy_width int, renderer *lipgloss.Renderer) Model {
	m := Model{
		row:        row,
		col:        col,
		r:          renderer,
		phy_width:  phy_width,
		curPos:     pos{x: 0, y: 0},
		KeyMap:     DefaultKeyMap(),
		mine_count: (row * col) * MINE_PERCENTAGE / 100,
		gameState:  starting,
		flags:      make(map[int]interface{}),
		help:       help.New(),
	}
	m.styles = m.makeStyle(m.r)
	m.init()
	m.help.ShowAll = true

	return m
}

func (m *Model) UpdatePhyWidth(w int) {
	m.phy_width = w
	m.help.Width = w
}

func (m *Model) choseMines(p pos) []int {
	curPosIdx := p.x*m.col + p.y // i*col + j: row major order

	all := make([]int, m.row*m.col-1)
	choosen := []int{}

	c := 0
	for i := 0; i < m.row*m.col; i++ { // TODO: simplefiy logic. User selection should not be mine
		if i != curPosIdx {
			all[c] = i
			c++
		}
	}

	for i := 0; i < m.mine_count; i++ {
		r := rand.IntN(len(all))
		choosen = append(choosen, all[r])

		// remove all[r]
		all = append(all[:r], all[r+1:]...)
	}

	return choosen
}

// idx = i*col + j
func idxToPos(idx, row, col int) (int, int) {
	i := idx / col
	j := idx - i*col

	return i, j
}

func posToIdx(p pos, row, col int) int {
	return p.x*col + p.y
}

func (m *Model) countMine(x, y int) int {
	c := 0
	board := m.board

	for i := max(0, x-1); i <= min(m.row-1, x+1); i++ {
		for j := max(0, y-1); j <= min(m.col-1, y+1); j++ {
			if !(i == x && j == y) && board[i][j].mine {
				c++
			}
		}
	}

	return c
}

func (m *Model) init() {
	m.board = make([][]cell, m.row)

	// initilize empty board
	for i := 0; i < m.row; i++ {
		m.board[i] = make([]cell, m.col)
	}
}

// called when the user, clicks on the firt cells
// makes the board
func (m *Model) generateMines() {
	log.Println("Generating Mines")

	mineIdxs := m.choseMines(m.curPos)

	for _, idx := range mineIdxs {
		i, j := idxToPos(idx, m.row, m.col)
		m.board[i][j].mine = true
	}

	for i := 0; i < m.row; i++ {
		for j := 0; j < m.col; j++ {
			m.board[i][j].mine_count = m.countMine(i, j)
		}
	}
}

func (m Model) hasWon() bool {
	if m.gameState == out {
		return false
	}

	for _, row := range m.board {
		for _, cell := range row {
			if cell.mine {
				if cell.shown {
					return false
				}
			} else {
				if !cell.shown {
					return false
				}
			}
		}
	}

	return true
}
