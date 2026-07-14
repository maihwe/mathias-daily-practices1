package ascii

import (
	"strings"
)

type ArtBuilder struct {
	text  string
	style string
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{
		style: "normal",
	}
}

func (a *ArtBuilder) AddText(text string) *ArtBuilder {
	a.text += text
	return a
}

func (a *ArtBuilder) SetStyle(style string) *ArtBuilder {
	switch style {
	case "normal", "bold", "italic", "outline":
		a.style = style
	default:
		panic("invalid style")
	}
	return a
}

func (a *ArtBuilder) Build() string {
	lines := make([]string, 8)

	var content string

	switch a.style {
	case "normal":
		content = a.text

	case "bold":
		for _, r := range a.text {
			content += string(r) + string(r)
		}

	case "italic":
		for i := range lines {
			lines[i] = strings.Repeat(" ", 7-i) + a.text
		}
		return strings.Join(lines, "\n")

	case "outline":
		content = "+" + strings.Repeat("-", len(a.text)) + "+"
		lines[0] = content
		for i := 1; i < 7; i++ {
			lines[i] = "|" + a.text + "|"
		}
		lines[7] = content
		return strings.Join(lines, "\n")
	}

	for i := range lines {
		lines[i] = content
	}
	return strings.Join(lines, "\n")

}
