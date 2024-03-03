package model

import (
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	browse "github.com/gjustoo/clib/cmd/browse"
)

type searchModel struct {
	prompt      string
	header      string
	fieldStyle  *lipgloss.Style
	headerStyle *lipgloss.Style
	// modelStyle  *lipgloss.Style
	answerField textinput.Model
}

func (s searchModel) Render() string {

	prompt := s.fieldStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left, s.prompt, s.answerField.View()))

	return s.headerStyle.Render(lipgloss.JoinVertical(lipgloss.Left, s.header, prompt))

}

func NewSearchModel() *searchModel {

	// header := `				 _____                    _
	// 		    / ____|                  | |
	// 			| |  __  ___   ___   __ _| | ___
	// 			| | |_ |/ _ \ / _ \ / _  | |/ _ \
	// 			| |__| | (_) | (_) | (_| | |  __/
	// 		     \_____|\___/ \___/ \__, |_|\___|
	// 								__/  |
	// 								|___/        `

	header := `	██████╗ ██████╗  █████╗ ██╗   ██╗███████╗    ███████╗███████╗ █████╗ ██████╗  ██████╗██╗  ██╗
	██╔══██╗██╔══██╗██╔══██╗██║   ██║██╔════╝    ██╔════╝██╔════╝██╔══██╗██╔══██╗██╔════╝██║  ██║
	██████╔╝██████╔╝███████║██║   ██║█████╗      ███████╗█████╗  ███████║██████╔╝██║     ███████║
	██╔══██╗██╔══██╗██╔══██║╚██╗ ██╔╝██╔══╝      ╚════██║██╔══╝  ██╔══██║██╔══██╗██║     ██╔══██║
	██████╔╝██║  ██║██║  ██║ ╚████╔╝ ███████╗    ███████║███████╗██║  ██║██║  ██║╚██████╗██║  ██║
	╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝    ╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝
																								 `

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA"))

	prompStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA"))

	s := &searchModel{prompt: "Search input: ", header: header, fieldStyle: &prompStyle, headerStyle: &headerStyle, answerField: textinput.New()}
	s.answerField.Focus()
	return s

}

func (m searchModel) Init() tea.Cmd {
	return nil
}

func (m searchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	log.Print("Triggered UPDATE from searchModel")

	switch msg := msg.(type) {
	case browse.HelloMsg:
		log.Print(" Entered helloMsg from searchModel")
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			log.Print("Entered ctrlc from searchModel")
			return m, tea.Quit
		case "enter":
			log.Print("entered enter keystroke from searchModel")
			return NewResultModel(), nil
		}
	}

	m.answerField, cmd = m.answerField.Update(msg)

	return m, cmd
}

func (m searchModel) View() string {
	// if m.width == 0 {
	// 	return "loading...."
	// }

	return lipgloss.JoinVertical(lipgloss.Left, m.Render())
}
