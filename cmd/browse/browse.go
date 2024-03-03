package cmd

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type HelloMsg string

func WaitASec() tea.Msg {
	time.Sleep(time.Second)
	return HelloMsg("Hi, there!")
}

// import (
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/gjustoo/clib/model/result"
// 	"github.com/gjustoo/clib/model/search"
// )

// const (
// 	SEARCH = iota + 1 // Search view
// 	RESULT            // Result view
// )

// func ModelHandler(r int) tea.Model {

// 	switch r {
// 	case SEARCH:
// 		return search.NewModel()
// 	case RESULT:
// 		return result.New()

// 	default:
// 		return nil
// 	}

// }
