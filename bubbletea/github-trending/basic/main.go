package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	repos []*Repo
	err   error
}

type errMsg struct{ error }

func (e errMsg) Error() string {
	return e.error.Error()
}

func newModel() model {
	return model{}
}

func main() {
	p := tea.NewProgram(newModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return fetchTrending
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	case []*Repo:
		m.repos = msg
		return m, nil

	default:
		return m, nil
	}
}

func (m model) View() string {
	var s string
	if m.err != nil {
		s = fmt.Sprintf("Fetch trending failed: %v", m.err)
	} else if len(m.repos) > 0 {
		for _, repo := range m.repos {
			s += repoText(repo)
		}
		s += "--------------------------------------"
	} else {
		s = " Fetching GitHub trending ..."
	}
	s += "\n\n"
	s += "Press q or ctrl + c or esc to exit..."
	return s + "\n"
}

func fetchTrending() tea.Msg {
	repos, err := getTrending("", "daily")
	if err != nil {
		return errMsg{err}
	}

	return repos
}

func repoText(repo *Repo) string {
	s := "--------------------------------------\n"
	s += fmt.Sprintf(`Repo:  %s | Language:  %s | Stars:  %d | Forks:  %d | Stars today:  %d
`, repo.Name, repo.Lang, repo.Stars, repo.Forks, repo.Add)
	s += fmt.Sprintf("Desc:  %s\n", repo.Desc)
	s += fmt.Sprintf("Link:  %s\n", repo.Link)
	return s
}
