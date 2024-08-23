package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/elliotcubit/gh-workflow-tui/pkg/ui"
)

func main() {
	client, err := api.DefaultRESTClient()
	if err != nil {
		os.Exit(1)
	}

	if _, err := tea.NewProgram(ui.New(client, nil), tea.WithAltScreen()).Run(); err != nil {
		os.Exit(1)
	}
}
