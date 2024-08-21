package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"

	"tsukuyomi/config"
)

func main() {
	styles := log.DefaultStyles()
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString(strings.ToUpper(log.DebugLevel.String())).
		Bold(true).
		Foreground(lipgloss.Color("63"))

	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString(fmt.Sprintf("%5s", strings.ToUpper(log.InfoLevel.String()))).
		Bold(true).
		Foreground(lipgloss.Color("27"))

	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString(fmt.Sprintf("%5s", strings.ToUpper(log.WarnLevel.String()))).
		Bold(true).
		Foreground(lipgloss.Color("120"))

	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString(strings.ToUpper(log.ErrorLevel.String())).
		Bold(true).
		Foreground(lipgloss.Color("204"))

	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString(strings.ToUpper(log.FatalLevel.String())).
		Bold(true).
		Foreground(lipgloss.Color("92"))

	prefixStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("10"))

	log.SetStyles(styles)
	log.SetReportCaller(true)
	log.SetReportTimestamp(true)
	log.SetPrefix(prefixStyle.Render("tsukuyomi"))
	log.SetTimeFormat("15:04")
	log.SetLevel(log.DebugLevel)

	log.Info("Starting...")
	config := config.Load()

	log.Info("Config loaded!", "config", config)
}
