package color

import "fmt"

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

func textColor(count int) string {
	if count == NONE {
		return fmt.Sprintf("%s[%dm", reset, count)
	}

	return fmt.Sprintf("%s[3%dm", reset, count)
}

func Color(count int, text string) string {
	return textColor(count) + text + textColor(NONE)
}
