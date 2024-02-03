package micropomo

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func info(logstring string) {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", logstring)
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}
}
