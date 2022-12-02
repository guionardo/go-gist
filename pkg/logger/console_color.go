package logger

import "runtime"

const (
	Reset int = iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	Gray
	White
)

var Colors = NewConsoleColor()

type ConsoleColor struct {
	colors map[int]string
}

func NewConsoleColor() (colors ConsoleColor) {
	colors = ConsoleColor{
		colors: make(map[int]string),
	}

	if runtime.GOOS == "windows" {
		return
	}
	colors.colors[Reset] = "\033[0m"
	colors.colors[Red] = "\033[31m"
	colors.colors[Green] = "\033[32m"
	colors.colors[Yellow] = "\033[33m"
	colors.colors[Blue] = "\033[34m"
	colors.colors[Purple] = "\033[35m"
	colors.colors[Cyan] = "\033[36m"
	colors.colors[Gray] = "\033[37m"
	colors.colors[White] = "\033[97m"

	return
}

func (c *ConsoleColor) Paint(text string, color int) string {
	if color, ok := c.colors[color]; ok {
		return color + text + c.colors[Reset]
	}
	return text
}
