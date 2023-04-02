package tui

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/aldernero/gaul"
	"github.com/aldernero/glitch-clock/pkg/glyph"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rivo/uniseg"
)

const (
	timeout              = 365 * 24 * time.Hour
	timeHorizontalOffset = 16
	maxDateWidth         = 86
	glitchChance         = 0.75
)

var glyphSet *glyph.GlyphSet

type Model struct {
	timer         timer.Model
	showDate      bool
	useLocalTime  bool
	dateSeparator string
}

func StartTea(date, local bool, sep string) {
	glyphSet = glyph.NewGlyphSet()
	m := Model{
		timer:         timer.NewWithInterval(timeout, 250*time.Millisecond),
		showDate:      date,
		useLocalTime:  local,
		dateSeparator: sep,
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func (m Model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}
	m.timer, cmd = m.timer.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var output string
	var gradient gaul.Gradient
	now := time.Now()
	var t, d string
	dateFormat := fmt.Sprintf("2006%s01%s02", m.dateSeparator, m.dateSeparator)
	if m.useLocalTime {
		t = now.Format("15:04:05")
		d = now.Format(dateFormat)
	} else {
		t = now.UTC().Format("15:04:05")
		d = now.UTC().Format(dateFormat)
	}
	if m.showDate {
		timeStr := glyphSet.Glyphify(t)
		dateStr := glyphSet.Glyphify(d)
		output = lipgloss.JoinVertical(0.0, timeStyle(timeStr), dateStyle(dateStr))
		gradient = synthWaveGradientFull
	} else {
		output = glyphSet.Glyphify(t)
		gradient = synthWaveGradientHalf
	}
	var w, h, index int
	var whitespace = []int{}
	var chars = []int{}
	for _, line := range strings.Split(output, "\n") {
		c := uniseg.GraphemeClusterCount(line)
		if c > w {
			w = c
		}
		for _, r := range line {
			if r == ' ' || r == '\t' {
				whitespace = append(whitespace, index)
			} else {
				chars = append(chars, index)
			}
			index++
		}
		h++
	}
	glitchStr := applyGlitches(output, whitespace, chars)
	palette := gradient.LinearPaletteStrings(h)
	var view string
	for i, line := range strings.Split(glitchStr, "\n") {
		for _, r := range line {
			if r == ' ' || r == '\t' {
				view += string(r)
			} else {
				c := lipgloss.Color(palette[i])
				view += lipgloss.NewStyle().Foreground(c).Render(string(r))
			}
		}
		view += "\n"
	}
	return viewStyle(view)
}

func applyGlitches(s string, whitespace []int, chars []int) string {
	if rand.Float32() <= glitchChance {
		return s
	}
	output := []rune(strings.Clone(s))
	swaps := rand.Intn(5) + 1
	m := len(whitespace)
	n := len(chars)
	for i := 0; i < swaps; i++ {
		var a, b int
		if rand.Float32() > 0.5 {
			a = whitespace[rand.Intn(m)]
			b = chars[rand.Intn(n)]
		} else {
			a = chars[rand.Intn(n)]
			b = chars[rand.Intn(n)]
		}
		output[a], output[b] = output[b], output[a]
	}
	return string(output)
}
