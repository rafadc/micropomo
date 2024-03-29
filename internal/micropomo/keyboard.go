package micropomo

import (
	tea "github.com/charmbracelet/bubbletea"
)

func manageKeys(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc", "ctrl+c":
		return m, tea.Quit
	case "r":
		return m.resetClock(), nil
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
