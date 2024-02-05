package micropomo

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	elapsedTime uint
	maxTime     uint
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

func InitialModel(minutes uint) model {
	return model{
		elapsedTime: 0,
		maxTime:     minutes * 60,
		clockStatus: Stopped,
		progress:    progress.New(gradient),
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
		return manageKeys(m, msg)

	case TickMsg:
		return m.tick()

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func (m model) resetClock() tea.Model {
	m.elapsedTime = 0
	m.clockStatus = Stopped
	m.progress = progress.New(gradient)
	return m
}

func (m model) tick() (tea.Model, tea.Cmd) {
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
		if m.elapsedTime >= m.maxTime+10 {
			return m.resetClock(), tickEvery()
		}
	}
	return m, tickEvery()
}
