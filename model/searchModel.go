package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type searchModel struct {
	prompt      string
	header      string
	fieldStyle  *lipgloss.Style
	headerStyle *lipgloss.Style
	answerField textinput.Model
}

func (s searchModel) Render() string {
	prompt := s.fieldStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left, s.prompt, s.answerField.View()))
	return s.headerStyle.Render(lipgloss.JoinVertical(lipgloss.Left, s.header, prompt))
}

func NewSearchModel() *searchModel {

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

func (m searchModel) Init() tea.Cmd { return nil }

func (m searchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			r := GetResults(m.answerField.Value())
			return NewResultModel(m.answerField.Value(), r), nil
		}

	}

	m.answerField, cmd = m.answerField.Update(msg)

	return m, cmd
}

func (m searchModel) View() string { return lipgloss.JoinVertical(lipgloss.Left, m.Render()) }

func GetResults(query string) []Answer {

	query = url.QueryEscape(query)

	url := fmt.Sprintf("https://api.search.brave.com/res/v1/web/search?q=%s&count=20&result_filter=web", query)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Subscription-Token", "")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	m := map[string]interface{}{}

	err = json.Unmarshal(body, &m)

	if err != nil {
		panic(err)
	}

	r := parseMap(m)
	return r

}

func parseMap(aMap map[string]interface{}) []Answer {

	webMap, exists := aMap["web"]

	r := []Answer{}
	if exists {
		return parseMap(webMap.(map[string]interface{}))
	}

	webMap, exists = aMap["results"]

	if exists {

		for _, val := range webMap.([]interface{}) {
			r = append(r, parseResult(val.(map[string]interface{})))
		}
	}
	return r
}

func parseResult(rm map[string]interface{}) Answer {
	return Answer{title: rm["title"].(string), desc: cleanDesc(rm["description"].(string)), Url: rm["url"].(string)}
}

func cleanDesc(desc string) (a string) {

	a = strings.ReplaceAll(desc, "<strong>", "")
	a = strings.ReplaceAll(a, "</strong>", "")

	return

}
