package micropomo

import (
	tea "github.com/charmbracelet/bubbletea"
)

func manageKeys(msg tea.KeyMsg, m model) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc", "ctrl+c":
		return m, tea.Quit
	case "r":
		m.clockStatus = Stopped
		m.elapsedTime = 0
	case " ":
		switch m.clockStatus {
		case Running:
			m.clockStatus = Paused
		case Paused:
			m.clockStatus = Running
		case Stopped:
			m.clockStatus = Running
		}
	}

	return m, nil
}
