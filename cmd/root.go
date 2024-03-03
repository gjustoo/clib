/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gjustoo/clib/model"
)

func Execute() {

	m := model.NewSearchModel()

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal("err: %w", err)
	}

}
