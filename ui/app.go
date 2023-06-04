package ui

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/go-github/v52/github"
)

type AppModel struct {
	client *github.Client
	list   list.Model
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
	return tea.Batch(enqueuegetReposListCmd, m.list.StartSpinner())
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
	case getReposListMsg:
		log.Println("getReposListCmd")
		cmds = append(cmds, m.list.StartSpinner(), getReposCmd(m.client))
	case gotReposListMsg:
		log.Println("gotReposListCmd")
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
