package main

import (
	"context"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/go-github/v52/github"
	"github.com/urfave/cli/v2"
)

type AppModel struct {
	client *github.Client
}

func NewAppModal(client *github.Client) AppModel {
	return AppModel{client: client}
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch()
}

func (m AppModel) View() string {
	return ""
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Batch()
}

func main() {
	app := cli.NewApp()

	app.Authors = []*cli.Author{{
		Name: "thetnaingtn",
	}}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			EnvVars: []string{"GITHUB_TOKEN"},
			Name:    "token",
			Usage:   "Your Github Token",
			Aliases: []string{"t"},
		},
	}

	app.Action = func(c *cli.Context) error {

		token := c.String("token")
		if token == "" {
			return cli.Exit("missing github token", 1)
		}

		ctx := context.Background()

		client := github.NewTokenClient(ctx, token)

		m := NewAppModal(client)

		p := tea.NewProgram(m, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			return cli.Exit(err.Error(), 1)
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}

}
