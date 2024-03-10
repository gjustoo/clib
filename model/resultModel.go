package model

import (
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type resultModel struct {
	answers []Answer
	cursor  int
	query   string
	list    list.Model
}

type Answer struct {
	title string
	desc  string
	Url   string
}

func (i Answer) Title() string       { return i.title }
func (i Answer) Description() string { return i.desc }
func (i Answer) FilterValue() string { return i.desc }

func NewResultModel(query string, answers []Answer) *resultModel {

	items := []list.Item{}
	for _, a := range answers {

		items = append(items, a)

	}

	s := &resultModel{query: query + " : \n", cursor: 0, answers: answers, list: list.New(items, list.NewDefaultDelegate(), 50, 100)}
	s.list.Title = "Results of : " + query
	return s
}

func (m resultModel) Init() tea.Cmd {
	return nil
}

type tickMsg int

func tick() tea.Msg {
	time.Sleep(time.Second / 4)
	return tickMsg(1)
}
func (m resultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewSearchModel(), nil
		case "enter":
			open(m.answers[m.cursor].Url)
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case tickMsg:
		h, v := docStyle.GetFrameSize()
		tw, th, _ := term.GetSize(int(os.Stdout.Fd()))
		m.list.SetSize(th-h, tw-v)
		return m, tea.Batch(tick, func() tea.Msg { return tea.WindowSizeMsg{Width: h, Height: v} })
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m resultModel) View() string {

	return docStyle.Render(m.list.View())

}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
