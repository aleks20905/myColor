package myColor

import (
	"fmt"
)

const reset = "\033"

const (
	NONE = iota
	RED
	GREEN
	YELLOW
	BLUE
	PURPLE
	LITEBLUE
)

func textColor(color int) string {
	if color == NONE {
		return fmt.Sprintf("%s[%dm", reset, 0) // Reset should be "\033[0m"
	}
	return fmt.Sprintf("%s[3%dm", reset, color)
}

func Color(color int, text string) string {
	return textColor(color) + text + textColor(NONE)
}

func Red(text string) string {
	return Color(RED, text)
}

func Green(text string) string {
	return Color(GREEN, text)
}

func Yellow(text string) string {
	return Color(YELLOW, text)
}

func Blue(text string) string {
	return Color(BLUE, text)
}

func Purple(text string) string {
	return Color(PURPLE, text)
}

func LiteBlue(text string) string {
	return Color(LITEBLUE, text)
}
