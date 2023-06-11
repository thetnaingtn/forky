package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/go-github/v52/github"
	"github.com/thetnaingtn/forky/ui"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Authors = []*cli.Author{{
		Name:  "thetnaingtn",
		Email: "thetnaingtun.ucsy@gmail.com",
	}}
	app.Usage = "Synchronize your forks with ease."

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			EnvVars: []string{"GITHUB_TOKEN"},
			Name:    "token",
			Usage:   "Your Github Token",
			Aliases: []string{"t"},
		},
	}

	app.Action = func(c *cli.Context) error {

		log.SetFlags(0)
		f, err := tea.LogToFile(filepath.Join(os.TempDir(), "forky.log"), "")
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
		defer func() { _ = f.Close() }()

		token := c.String("token")
		if token == "" {
			return cli.Exit("missing github token", 1)
		}

		ctx := context.Background()

		client := github.NewTokenClient(ctx, token)

		m := ui.NewAppModel(client)

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
