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
	list.Styles.Title = listTitleStyle
	list.Title = "Forky"
	list.SetSpinner(spinner.MiniDot)

	list.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			keySelectToggle,
			keyMergeWithUpstream,
		}
	}

	list.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{keyRefresh}
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
		m.list.Title = "Getting forks. Hold tight!"
		m.list.SetItems([]list.Item{}) // reset to empty list!!
		m.list.SetShowStatusBar(false)
		m.list.SetShowHelp(false)
		cmds = append(cmds, m.list.StartSpinner(), getReposCmd(m.client))
	case gotReposListMsg:
		log.Println("gotReposListCmd")
		m.list.Title = "Forky"
		m.list.StopSpinner()
		m.list.SetShowStatusBar(true)
		m.list.SetShowHelp(true)
		cmds = append(cmds, m.list.SetItems(reposToItems(msg.repos)))
	case mergeSelectedReposMsg:
		m.list.Title = "Syncing with upstream repository!"
		items := m.list.Items()
		cmds = append(cmds, mergeReposCmd(m.client, items))
	case mergedSelectedReposMsg:
		m.list.Title = "Forky"
		m.list.StopSpinner()
		cmds = append(cmds, m.list.SetItems(msg.items))
	// key messages
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

		if key.Matches(msg, keySelectToggle) {
			cmds = append(cmds, m.toggleSelection())
		}

		if key.Matches(msg, keyMergeWithUpstream) {
			cmds = append(cmds, m.list.StartSpinner(), requestMergeReposCmd)
		}

		if key.Matches(msg, keyRefresh) {
			cmds = append(cmds, m.list.StartSpinner(), enqueuegetReposListCmd)
		}
	}

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m AppModel) toggleSelection() tea.Cmd {
	idx := m.list.Index()
	item := m.list.SelectedItem().(item)
	item.selected = !item.selected
	m.list.RemoveItem(idx)

	return m.list.InsertItem(idx, item)
}
