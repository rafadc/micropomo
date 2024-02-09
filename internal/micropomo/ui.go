package micropomo

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

var gradient = progress.WithScaledGradient("#FF4444", "#FF0000")

func (m model) View() string {
	if m.clockStatus == Finishing {
		style := lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63"))
		return style.Render("🎉  Time's up! 🎉")
	}
	s := fmt.Sprintf("%s ", m.statusIcon())
	s += fmt.Sprintf("%s - %s\n", formatTime(m.elapsedTime), formatTime(m.maxTime))
	s += m.progress.View()
	return s
}

func formatTime(t uint) string {
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
