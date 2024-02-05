package main

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/rafadc/micropomo/internal/micropomo"
)

func main() {
	minutes := uint(25)
	if len(os.Args) > 1 {
		min, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Invalid number of minutes for pomodoro: %v", err)
			os.Exit(1)
		}
		minutes = uint(min)
	}

	p := tea.NewProgram(micropomo.InitialModel(minutes))
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
