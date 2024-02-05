package micropomo

import (
	"fmt"
)

func (m model) View() string {
	s := fmt.Sprintf("%s ", m.statusIcon())
	s	+= fmt.Sprintf("%s - %s\n", formatTime(m.elapsedTime), formatTime(m.maxTime))
	s += m.progress.View()
	return s
}

func formatTime(t int) string {
	minutes := t / 60
	seconds := t % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func (m model) statusIcon() string {
	if m.clockStatus == Running {
		return "⏵"
	} else if m.clockStatus == Paused {
		return "⏸"
	} else {
		return "⏹"
	}
}
