package micropomo

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	elapsedTime int
	maxTime     int
	clockStatus Status
	progress    progress.Model
}

type Status int

const (
	Running Status = iota
	Paused
	Stopped
	Finishing
)

type TickMsg time.Time

func InitialModel() model {
	return model{
		elapsedTime: 0,
		maxTime:     25 * 60,
		clockStatus: Stopped,
		progress:    progress.New(progress.WithDefaultGradient()),
	}
}

func tickEvery() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return tickEvery()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "r":
			m = InitialModel()
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
	case TickMsg:
		switch m.clockStatus {
		case Running:
			m.elapsedTime++
			progressCmd := m.progress.SetPercent(float64(m.elapsedTime) / float64(m.maxTime))
			if m.elapsedTime >= m.maxTime {
				m.clockStatus = Finishing
			}
			return m, tea.Batch(tickEvery(), progressCmd)
		case Finishing:
			m.elapsedTime++
			if m.elapsedTime >= m.maxTime+5 {
				m.clockStatus = Stopped
			}
		}
		return m, tickEvery()

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("Elapsed time: %d - Max time: %d\n", m.elapsedTime, m.maxTime)
	s += m.progress.View()
	return s
}
