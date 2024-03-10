package model

import (
	"bytes"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	browse "github.com/gjustoo/clib/cmd/browse"
)

// import (
// 	"github.com/charmbracelet/bubbles/textinput"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

type resultModel struct {
	answers []Answer
	index   int
	query   string
	style   *lipgloss.Style
}

type Answer struct {
	Title       string
	Description string
	Url         string
}

func NewResultModel() *resultModel {

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA"))

	answers := []Answer{
		Answer{Title: "Como hacer las mejores patatas fritas del mundo ", Description: "Pues nada eso, que como hacer las mejores patatas fritas del mundo mundial", Url: "https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwiypdHG19iEAxWMV6QEHYDSDy4QFnoECBYQAQ&url=https%3A%2F%2Fwww.directoalpaladar.com%2Frecetario%2Fcomo-hacer-las-mejores-patatas-fritas-del-mundo-mundial&usg=AOvVaw3pI7jNA0O6YKlOpE8xHF7i&opi=89978449"},
		Answer{Title: "Como hacer las segundas mejores patatas fritas del mundo ", Description: "Pues nada eso, que como hacer las segundas mejores patatas fritas del mundo mundial", Url: "https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwiypdHG19iEAxWMV6QEHYDSDy4QFnoECBYQAQ&url=https%3A%2F%2Fwww.directoalpaladar.com%2Frecetario%2Fcomo-hacer-las-mejores-patatas-fritas-del-mundo-mundial&usg=AOvVaw3pI7jNA0O6YKlOpE8xHF7i&opi=89978449"},
		Answer{Title: "Como hacer las terceras mejores patatas fritas del mundo ", Description: "Pues nada eso, que como hacer las terceras mejores patatas fritas del mundo mundial", Url: "https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwiypdHG19iEAxWMV6QEHYDSDy4QFnoECBYQAQ&url=https%3A%2F%2Fwww.directoalpaladar.com%2Frecetario%2Fcomo-hacer-las-mejores-patatas-fritas-del-mundo-mundial&usg=AOvVaw3pI7jNA0O6YKlOpE8xHF7i&opi=89978449"},
		Answer{Title: "Como hacer las cuartas  mejores patatas fritas del mundo ", Description: "Pues nada eso, que como hacer las cuartas mejores patatas fritas del mundo mundial", Url: "https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwiypdHG19iEAxWMV6QEHYDSDy4QFnoECBYQAQ&url=https%3A%2F%2Fwww.directoalpaladar.com%2Frecetario%2Fcomo-hacer-las-mejores-patatas-fritas-del-mundo-mundial&usg=AOvVaw3pI7jNA0O6YKlOpE8xHF7i&opi=89978449"},
		Answer{Title: "Como hacer las quintas mejores patatas fritas del mundo ", Description: "Pues nada eso, que como hacer las quintas mejores patatas fritas del mundo mundial", Url: "https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwiypdHG19iEAxWMV6QEHYDSDy4QFnoECBYQAQ&url=https%3A%2F%2Fwww.directoalpaladar.com%2Frecetario%2Fcomo-hacer-las-mejores-patatas-fritas-del-mundo-mundial&usg=AOvVaw3pI7jNA0O6YKlOpE8xHF7i&opi=89978449"},
		Answer{Title: "Como hacer las sextas mejores patatas fritas del mundo ", Description: "Pues nada eso, que como hacer las sextas mejores patatas fritas del mundo mundial", Url: "https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwiypdHG19iEAxWMV6QEHYDSDy4QFnoECBYQAQ&url=https%3A%2F%2Fwww.directoalpaladar.com%2Frecetario%2Fcomo-hacer-las-mejores-patatas-fritas-del-mundo-mundial&usg=AOvVaw3pI7jNA0O6YKlOpE8xHF7i&opi=89978449"},
		Answer{Title: "Como hacer las septimas mejores patatas fritas del mundo ", Description: "Pues nada eso, que como hacer las septimas mejores patatas fritas del mundo mundial", Url: "https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwiypdHG19iEAxWMV6QEHYDSDy4QFnoECBYQAQ&url=https%3A%2F%2Fwww.directoalpaladar.com%2Frecetario%2Fcomo-hacer-las-mejores-patatas-fritas-del-mundo-mundial&usg=AOvVaw3pI7jNA0O6YKlOpE8xHF7i&opi=89978449"},
	}

	s := &resultModel{query: "Patatas fritas", index: 0, style: &style, answers: answers}
	return s
}

func (m resultModel) Init() tea.Cmd {
	return nil
}

func (m resultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	log.Print("Triggered UPDATE from resultModel")

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			log.Print("Entered ctrlc from resultModel")
			return m, tea.Quit
		case "esc":
			log.Print("Entered ctrlc from resultmodel")
			return NewSearchModel(), browse.WaitASec

		}
	}

	return m, nil
}

func (m resultModel) View() string {

	return m.style.Render(lipgloss.JoinVertical(lipgloss.Left, m.query, " Results : ", answToString(m.answers)))
}

func answToString(a []Answer) string {

	var out bytes.Buffer

	for _, ans := range a {

		out.WriteString(ans.string())
	}

	return out.String()
}

func (a Answer) string() string {

	var out bytes.Buffer

	result := fmt.Sprintf("%s \n %s \n %s \n", a.Title, a.Description, a.Url)

	out.WriteString(result)

	return out.String()

}
