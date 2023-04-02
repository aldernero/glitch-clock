package glyph

import (
	"fmt"
	"github.com/rivo/uniseg"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var RuneMap = map[rune]string{
	'0': `
 ▞████▚ 
▟█  ███▙
▓▓ ▟▛ ▓▓
▒▒▒▒  ▒▒
 ░░░░░░`,
	'1': `
 ▞▚
▟██
 ▓▓
 ▒▒
 ░░`,
	'2': `
▞████▚ 
     ██ 
 ▓▓▓▓▓  
▒▒      
░░░░░░░`,
	'3': `
▞████▚  
     ██
 ▓▓▓▓▓  
     ▒▒
░░░░░░`,
	'4': `
▞▚   ▞▚ 
██   ██ 
▓▓▓▓▓▓▓ 
     ▒▒ 
     ░░`,
	'5': `
▞█████▚ 
██      
▓▓▓▓▓▓▓ 
     ▒▒ 
░░░░░░░`,
	'6': `
 ▞████▚  
▟█       
▓▓▓▓▓▓▓  
▒▒    ▒▒ 
 ░░░░░░`,
	'7': `
▞█████▚ 
     █▙
    ▓▓  
   ▒▒   
   ░░`,
	'8': `
 ▞███▚  
▟█   █▙
 ▓▓▓▓▓  
▒▒   ▒▒
 ░░░░░ `,
	'9': `
 ▞███▚ 
▟█   █▙
 ▓▓▓▓▓▓
     ▒▒
 ░░░░░ `,
	':': `
█▚

▒▒`,
	'-': `

▓▓▓▓▓

`,
	'/': `
    █▞
   █▛
  ▓▓
 ▒▒
░░`,
}

var alphabet = []rune("0123456789:-/")

type Glyph struct {
	text          string
	width, height int
}

func NewGlyph(text string) Glyph {
	var w, h int
	txt := strings.TrimRight(text, "\n")
	txt = strings.TrimLeft(txt, "\n")
	for _, line := range strings.Split(txt, "\n") {
		count := uniseg.GraphemeClusterCount(line)
		if count > w {
			w = count
		}
		h++
	}
	return Glyph{
		text:   txt,
		width:  w,
		height: h,
	}
}

func (g *Glyph) Render(style lipgloss.Style) string {
	var output string
	glyphStyle := lipgloss.NewStyle().AlignHorizontal(lipgloss.Left).Width(g.width)
	for _, line := range strings.Split(g.text, "\n") {
		var txt string
		if line != "" {
			txt = glyphStyle.Render(line)
		}
		output += txt + "\n"
	}
	return style.Render(output)
}

type GlyphSet struct {
	runeMap map[rune]Glyph
}

func NewGlyphSet() *GlyphSet {
	glyphs := make(map[rune]Glyph)
	for _, r := range alphabet {
		text, ok := RuneMap[r]
		if !ok {
			panic(fmt.Sprintf("no definition for rune %v", r))
		}
		glyphs[r] = NewGlyph(text)
	}
	return &GlyphSet{
		runeMap: glyphs,
	}
}

func (g *GlyphSet) Glyphify(msg string) string {
	var output string
	var glyphs []Glyph
	var maxHeight int
	for _, r := range msg {
		glyph, ok := g.runeMap[r]
		if !ok {
			continue
		}
		if glyph.height > maxHeight {
			maxHeight = glyph.height
		}
		glyphs = append(glyphs, glyph)
	}
	for _, glyph := range glyphs {
		style := glyphsStyle.Height(maxHeight + 1)
		output = lipgloss.JoinHorizontal(lipgloss.Top, output, glyph.Render(style))
	}
	return output
}
