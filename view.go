package main

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	appNameStyle = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)

	faintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)

	enumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
)

func (m model) View() string {

	s := appNameStyle.Render("NOTES APP") + "\n\n"

	if m.state == titleView {
		s += "Note title:\n\n"
		s += m.textinput.View() + "\n\n"
		s += faintStyle.Render("enter - save, esq - discard")
	}

	if m.state == bodyView {
		s += "Note: " + m.currNote.Title + "\n\n"
		s += m.textarea.View() + "\n\n"
		s += faintStyle.Render("ctrl+s - save, esq - discard")
	}

	if m.state == listView {
		for i, n := range m.notes {
			prefix := ""
			if i == m.listIndex {
				prefix = "-> "
			}
			shortBody := strings.ReplaceAll(n.Body, "\n", " ")
			if len(shortBody) > 30 {
				shortBody = shortBody[:30] + "..."
			}
			s += enumeratorStyle.Render(prefix) + n.Title + " | " + faintStyle.Render(shortBody) + "\n\n"
		}
		s += faintStyle.Render("n - new note, d - delete note, q - quit, up/down/j/k to navigate, enter to edit") + "\n"
	}

	return s
}
