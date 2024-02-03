package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/rafadc/micropomo/internal/micropomo"
)

func main() {
	p := tea.NewProgram(micropomo.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
