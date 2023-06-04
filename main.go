package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/go-github/v52/github"
	"github.com/urfave/cli/v2"
)

var (
	keySelectToggle       = key.NewBinding(key.WithKeys(" "), key.WithHelp("space", "toggle selected item"))
	keyMergedWithUpStream = key.NewBinding(key.WithKeys("m"), key.WithHelp("m", "merge selected item with upstream branch"))
)

var (
	errorColor = lipgloss.AdaptiveColor{
		Light: "#e94560",
		Dark:  "#f05945",
	}
	listStyle    = lipgloss.NewStyle().Margin(2)
	detailsStyle = lipgloss.NewStyle().PaddingLeft(2)

	errorStyle = lipgloss.NewStyle().Foreground(errorColor)
)

func enqueueGetReposCmd() tea.Msg {
	return GetReposCmd{}
}

func getAllRepositories(client *github.Client) tea.Cmd {
	return func() tea.Msg {
		var repos []*github.Repository
		opt := &github.RepositoryListOptions{Type: "all", Sort: "updated", Direction: "desc"}
		r, _, err := client.Repositories.List(context.Background(), "thetnaingtn", opt)
		if err != nil {
			log.Println(err)
		}
		for _, repo := range r {
			if !repo.GetFork() {
				continue
			}
			repos = append(repos, repo)
		}
		return GotReposCmd{repos: repos}
	}
}

func reposToItems(repos []*github.Repository) []list.Item {
	items := make([]list.Item, 0, len(repos))
	for _, repo := range repos {
		items = append(items, Item{repo: repo})
	}

	return items
}

type AppModel struct {
	client *github.Client
	list   list.Model
}

type GetReposCmd struct{}
type GotReposCmd struct {
	repos []*github.Repository
}

type Item struct {
	repo     *github.Repository
	selected bool
}

func (i Item) Title() string {
	return i.repo.GetFullName()
}

func (i Item) Description() string {
	return i.repo.GetDescription()
}

func (i Item) FilterValue() string {
	return "  " + i.repo.GetName()
}

func NewAppModel(client *github.Client) AppModel {
	list := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Forky"
	list.SetSpinner(spinner.MiniDot)

	list.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			keySelectToggle,
			keyMergedWithUpStream,
		}
	}

	return AppModel{client: client, list: list}
}

func (m AppModel) Init() tea.Cmd {
	return tea.Batch(enqueueGetReposCmd, m.list.StartSpinner())
}

func (m AppModel) View() string {
	return m.list.View()
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		log.Println("tea.WindowSizeMsg")
		top, right, bottom, left := listStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	case GetReposCmd:
		log.Println("GetReposCmd")
		cmds = append(cmds, m.list.StartSpinner(), getAllRepositories(m.client))
	case GotReposCmd:
		m.list.StopSpinner()
		cmds = append(cmds, m.list.SetItems(reposToItems(msg.repos)))
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
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

		m := NewAppModel(client)

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
