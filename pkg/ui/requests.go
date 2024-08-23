package ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cli/go-gh/v2/pkg/api"
)

type workflow struct {
	Id   int
	Name string
}

func (i workflow) Title() string       { return i.Name }
func (i workflow) Description() string { return "" }
func (i workflow) FilterValue() string { return i.Name }

type Response struct {
	Workflows []workflow
}

var linkRE = regexp.MustCompile(`<([^>]+)>;\s*rel="([^"]+)"`)

func findNextPage(response *http.Response) string {
	for _, m := range linkRE.FindAllStringSubmatch(response.Header.Get("Link"), -1) {
		if len(m) > 2 && m[2] == "next" {
			return m[1]
		}
	}
	return ""
}

func (m model) getPage() tea.Msg {
	if m.nextPage == "" {
		return nil
	}

	response, err := m.client.Request(http.MethodGet, m.nextPage, nil)
	if err != nil {
		return nil
	}
	var r Response
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&r)
	if err != nil {
		return nil
	}
	if err := response.Body.Close(); err != nil {
		return nil
	}
	nextPage := findNextPage(response)

	return resultPage{
		workflows: r.Workflows,
		nextPage:  nextPage,
	}
}

func dispatchWorkflow(client *api.RESTClient, w workflow) tea.Cmd {
	return func() tea.Msg {
		url := "repos/" + repo() + "/actions/workflows/" + strconv.Itoa(w.Id) + "/dispatches"
		body := []byte(`{"ref":"main"}`)
		client.Request(http.MethodPost, url, bytes.NewReader(body))
		return statusMessage{
			status: fmt.Sprintf("Kicked off workflow run for %s", w.Name),
		}

	}
}
