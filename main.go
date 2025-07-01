package main

import (
	"github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	m := NewModel()

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Unable to run program: %v", err)
	}
}
