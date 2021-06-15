package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	cyan   = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF"))
	green  = lipgloss.NewStyle().Foreground(lipgloss.Color("#32CD32"))
	gray   = lipgloss.NewStyle().Foreground(lipgloss.Color("#696969"))
	gold   = lipgloss.NewStyle().Foreground(lipgloss.Color("#B8860B"))
	purple = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF"))
)

const (
	CountPerPage = 5
)

type model struct {
	repos     []*Repo
	err       error
	curPage   int
	totalPage int
	spinner   spinner.Model
}

type errMsg struct{ error }

func (e errMsg) Error() string {
	return e.error.Error()
}

func newModel() model {
	sp := spinner.NewModel()
	sp.Style = purple

	return model{
		spinner: sp,
	}
}

func main() {
	p := tea.NewProgram(newModel())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		spinner.Tick,
		fetchTrending,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "pgdown":
			if m.curPage < m.totalPage-1 {
				m.curPage++
			}
			return m, nil
		case "pgup":
			if m.curPage > 0 {
				m.curPage--
			}
			return m, nil
		default:
			return m, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case errMsg:
		m.err = msg
		return m, nil

	case []*Repo:
		m.repos = msg
		m.totalPage = (len(msg) + CountPerPage - 1) / CountPerPage
		return m, nil

	default:
		return m, nil
	}
}

func (m model) View() string {
	var s string
	if m.err != nil {
		s = gold.Render(fmt.Sprintf("fetch trending failed: %v", m.err))
	} else if len(m.repos) > 0 {
		start, end := m.curPage*CountPerPage, (m.curPage+1)*CountPerPage
		if end > len(m.repos) {
			end = len(m.repos)
		}
		for _, repo := range m.repos[start:end] {
			s += repoText(repo)
		}
		s += cyan.Render("--------------------------------------")
	} else {
		s = m.spinner.View() + gold.Render(" Fetching GitHub trending ...")
	}
	s += "\n\n"
	if m.totalPage > 1 {
		s += gray.Render("Pagedown to next page, pageup to prev page.")
		s += "\n"
	}
	s += gray.Render("Press q or ctrl + c or esc to exit...")
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
	s := cyan.Render("--------------------------------------") + "\n"
	s += fmt.Sprintf(`Repo:  %s | Language:  %s | Stars:  %s | Forks:  %s | Stars today:  %s
`, cyan.Render(repo.Name), cyan.Render(repo.Lang), cyan.Render(strconv.Itoa(repo.Stars)),
		cyan.Render(strconv.Itoa(repo.Forks)), cyan.Render(strconv.Itoa(repo.AddStars)))
	s += fmt.Sprintf("Desc:  %s\n", green.Render(repo.Desc))
	s += fmt.Sprintf("Link:  %s\n", gray.Render(repo.Link))
	return s
}
