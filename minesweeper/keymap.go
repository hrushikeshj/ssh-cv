package minesweeper

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Down  key.Binding
	Up    key.Binding
	Left  key.Binding
	Right key.Binding
	Enter key.Binding
	Flag key.Binding
	Restart key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Flag, k.Restart}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Enter}, {k.Flag}, {k.Restart},
		{k.Up},{k.Down}, {k.Left}, {k.Right}, // first column
	}
}

func DefaultKeyMap() keyMap {
	return keyMap{
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "down"),
		),
		Left: key.NewBinding(
			key.WithKeys("left"),
			key.WithHelp("←", "left"),
		),
		Right: key.NewBinding(
			key.WithKeys("right"),
			key.WithHelp("→", "right"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("↵", "Reveal"),
		),
		Flag: key.NewBinding(
			key.WithKeys("f", "F"),
			key.WithHelp("f", "Flag"),
		),
		Restart: key.NewBinding(
			key.WithKeys("r", "R"),
			key.WithHelp("r", "Restart"),
		),
	}
}
