package tui

import (
	"github.com/aldernero/gaul"
	"github.com/charmbracelet/lipgloss"
)

var viewStyle = lipgloss.NewStyle().Margin(1, 0, 0, 0).Render
var dateStyle = lipgloss.NewStyle().Width(maxDateWidth).AlignHorizontal(lipgloss.Center).Render
var timeStyle = lipgloss.NewStyle().Width(maxDateWidth).MarginLeft(timeHorizontalOffset).Render

var synthWaveGradientFull = gaul.NewGradientFromNamed([]string{"cyan", "yellow", "magenta"})
var synthWaveGradientHalf = gaul.NewGradientFromNamed([]string{"yellow", "magenta"})
