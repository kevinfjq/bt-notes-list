package main

import "github.com/charmbracelet/bubbletea"

const (
	listView uint = iota
	titleView
	bodyView
)

type model struct {
	state uint
	//store Store
	//textarea.Model
	// ... other fields as needed
}

func NewModel() model {
	return model{
		state: listView,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
