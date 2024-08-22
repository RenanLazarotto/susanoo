package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func SetupLog() {
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

	logger := log.New(os.Stdout)
	logger.SetStyles(styles)
	logger.SetReportCaller(true)
	logger.SetReportTimestamp(true)
	logger.SetPrefix(prefixStyle.Render("tsukuyomi"))
	logger.SetTimeFormat("15:04")
	logger.SetLevel(log.DebugLevel)

	log.SetDefault(logger)
}
