package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hrushikeshj/ssh-cv/cv"
	"os"
)

func main() {
	m := cv.NewModel()
	// fmt.Println(m.RenderCV())

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
