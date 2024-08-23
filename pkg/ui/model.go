package ui

import (
	"slices"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/repository"
)

type model struct {
	client   *api.RESTClient
	nextPage string

	list         list.Model
	delegateKeys *delegateKeyMap
}

func repo() string {
	s, err := repository.Current()
	if err != nil {
		return ""
	}
	return s.Owner + "/" + s.Name
}

func New(client *api.RESTClient, workflows []workflow) model {
	var (
		delegateKeys = newDelegateKeyMap()
	)

	slices.SortFunc(workflows, func(a, b workflow) int {
		return strings.Compare(a.Name, b.Name)
	})

	items := make([]list.Item, len(workflows))
	for i, v := range workflows {
		items[i] = v
	}

	// Setup list
	delegate := newItemDelegate(client, delegateKeys)
	list := list.New(items, delegate, 0, 0)
	list.Title = "Github Workflows"
	list.StatusMessageLifetime = time.Second * 10
	list.Styles.Title = titleStyle

	return model{
		client:       client,
		nextPage:     "repos/" + repo() + "/actions/workflows?per_page=100",
		list:         list,
		delegateKeys: delegateKeys,
	}
}

func (m model) Init() tea.Cmd {
	return m.getPage
}

type resultPage struct {
	workflows []workflow
	nextPage  string
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case resultPage:
		items := m.list.Items()
		for _, v := range msg.workflows {
			items = append(items, v)
		}
		m.nextPage = msg.nextPage
		cmds = append(cmds, m.list.SetItems(items), m.getPage)
	}

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return appStyle.Render(m.list.View())
}
