package ascii

import (
	"fmt"
	"strings"
)

type Animation struct {
	text   string
	frames int
	data   []string
}

func NewAnimation(text string, frames int) *Animation {
	return &Animation{
		text:   text,
		frames: frames,
		data:   make([]string, frames),
	}
}

func MakeFrame(content string) string {
	lines := make([]string, 10)

	for i := range lines {
		lines[i] = content
	}
	return strings.Join(lines, "\n")
}

func (a *Animation) GenerateSpinFrames() {
	for i := 0; i < a.frames; i++ {
		prefix := strings.Repeat(" ", i)
		a.data[i] = MakeFrame(prefix + a.text)
	}
}

func (a *Animation) GenerateWaveFrames() {
	for i := 0; i < a.frames; i++ {
		lines := make([]string, 10)
		for j := 0; j < 10; j++ {
			offset := (i + j) % 5

			lines[j] = strings.Repeat(" ", offset) + a.text
		}
		a.data[i] = strings.Join(lines, "\n")
	}
}

func (a *Animation) GenerateZoomFrames() {
	for i := 0; i < a.frames; i++ {
		scale := i + 1

		var zoom string

		for _, r := range a.text {
			zoom += strings.Repeat(string(r), scale)
		}
		a.data[i] = MakeFrame(zoom)
	}
}

func (a *Animation) GetFrame(index int) string {
	if len(a.data) == 0 {
		return " "
	}
	return  a.data[index%len(a.data)]
}

func (a *Animation) Play() string {
	var result strings.Builder

	for i, frame := range a.data {
		result.WriteString(
			fmt.Sprintf("=== frame %d ===\n%s\n\n", i+1, frame),
		)
	}
	return result.String()
}