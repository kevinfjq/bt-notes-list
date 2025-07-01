package main

import (
	"github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	store := &Store{}
	if err := store.Init(); err != nil {
		log.Fatalf("Unable to initialize store: %v", err)
	}
	m := NewModel(store)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Unable to run program: %v", err)
	}
}
