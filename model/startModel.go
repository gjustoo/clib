package model

// import (
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// 	searchinput "github.com/gjustoo/clib/model/searchInput"
// )

// type Startmodel struct {
// 	text   string
// 	width  int
// 	height int
// 	style  *lipgloss.Style
// }

// func New() *Startmodel {

// 	return &Startmodel{text: text, style: &s}
// }
// func (s Startmodel) Render() string {

// 	return s.style.Render(s.text)

// }

// func (m Startmodel) Init() tea.Cmd {
// 	return nil
// }

// func (m Startmodel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		m.width = msg.Width
// 		m.height = msg.Height
// 		return searchinput.NewSearchInput("Google search "), nil
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c":
// 			return m, tea.Quit
// 		}
// 	}

// 	return m, nil
// }

// func (m Startmodel) View() string {
// 	if m.width == 0 {
// 		return "loading...."
// 	}

// 	return lipgloss.JoinVertical(lipgloss.Center, m.Render())
// }
