package main

import (
	"fmt"
	"os"
	"raven/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	ui.Splash()
	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
