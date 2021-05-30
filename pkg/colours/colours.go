package colours

import (
	"fmt"
	"strconv"
)

var esc = "\033"
var reset = esc + "[0m"

func printWithCode(s string, c int) string {
	return fmt.Sprintf(esc+"["+strconv.Itoa(c)+"m%s"+reset, s)
}

// Grey formats colored terminal text.
func Grey(s string) string {
	return printWithCode(s, 2)
}

// Underline formats underlined terminal text.
func Underline(s string) string {
	return printWithCode(s, 4)
}

// Red formats colored terminal text.
func Red(s string) string {
	return printWithCode(s, 31)
}

// Yellow formats colored terminal text.
func Yellow(s string) string {
	return printWithCode(s, 33)
}

// Blue formats colored terminal text.
func Blue(s string) string {
	return printWithCode(s, 34)
}
